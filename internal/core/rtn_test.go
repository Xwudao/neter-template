package core

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/Xwudao/neter-template/internal/domain/errs"
	"github.com/Xwudao/neter-template/internal/validate"
)

func TestNewRtnWithErrUsesBusinessMessage(t *testing.T) {
	status := NewRtnWithErr(errs.NewConflict("名称已存在", errors.New("unique index")))

	if status.Code != CodeError || status.Message != "名称已存在" {
		t.Fatalf("unexpected status: %#v", status)
	}
}

func TestErrToRtnUsesRtnStatusError(t *testing.T) {
	status := errToRtn(NewRtnError(CodeNeedVIP, "需要 VIP"))

	if status == nil || status.Code != CodeNeedVIP || status.Message != "需要 VIP" {
		t.Fatalf("unexpected status: %#v", status)
	}
}

func TestErrToRtnFallsBackToNewRtnWithErr(t *testing.T) {
	status := errToRtn(errors.New("boom"))

	if status == nil || status.Code != CodeError || status.Message != "boom" {
		t.Fatalf("unexpected status: %#v", status)
	}
}

// optimizeTracker 记录是否被自动调用。
type optimizeTracker struct {
	called bool
}

func (o *optimizeTracker) Optimize() {
	o.called = true
}

type optimizeErrTracker struct {
	called bool
	err    error
}

func (o *optimizeErrTracker) Optimize() error {
	o.called = true
	return o.err
}

func TestOptimizeRequestCallsNoArgOptimize(t *testing.T) {
	req := &optimizeTracker{}
	if !optimizeRequest(req) {
		t.Fatal("expected optimize success")
	}
	if !req.called {
		t.Fatal("expected Optimize() to be called")
	}
}

func TestOptimizeRequestCallsErrorOptimize(t *testing.T) {
	req := &optimizeErrTracker{}
	if !optimizeRequest(req) {
		t.Fatal("expected optimize success")
	}
	if !req.called {
		t.Fatal("expected Optimize() to be called")
	}
}

func TestOptimizeRequestFailsOnError(t *testing.T) {
	req := &optimizeErrTracker{err: errors.New("bad params")}
	if optimizeRequest(req) {
		t.Fatal("expected optimize failure")
	}
	if !req.called {
		t.Fatal("expected Optimize() to be called")
	}
}

func TestOptimizeRequestSkipsNonOptimizer(t *testing.T) {
	type plain struct {
		X int
	}
	if !optimizeRequest(&plain{X: 1}) {
		t.Fatal("expected optimize success for non-optimizer")
	}
}

// messageValidator 提供代码式字段校验。
type messageValidator struct {
	Value string `json:"value"`
}

func (m *messageValidator) Validate() error {
	return validate.Validate(
		validate.Field("value", m.Value,
			validate.Message("值不能为空", validate.Required()),
		),
	)
}

func TestDefaultBindErrorMapperKeepsSafeMapping(t *testing.T) {
	err := defaultBindErrorMapper(&messageValidator{}, errors.New("unused"))
	// 绑定解析失败统一走“参数错误”安全映射，不泄漏原始错误
	if err == nil || err.Error() != "参数错误" {
		t.Fatalf("unexpected err: %v", err)
	}
}

// validatingOptimizer 同时实现 Validate 与 Optimize，用于断言调用顺序。
type validatingOptimizer struct {
	Value     string `json:"value"`
	optimized bool
}

func (v *validatingOptimizer) Validate() error {
	if v.Value == "" {
		return validate.Validate(
			validate.Field("value", v.Value,
				validate.Message("值不能为空", validate.Required()),
			),
		)
	}
	return nil
}

func (v *validatingOptimizer) Optimize() {
	v.optimized = true
}

// TestBindAndRespondValidatesBeforeOptimize 断言：绑定成功后 Validate 先于 Optimize 调用；
// 校验失败时返回原包装结构 + 第一条中文消息，且不执行 Optimize。
func TestBindAndRespondValidatesBeforeOptimize(t *testing.T) {
	gin.SetMode(gin.TestMode)

	handler := func(c *gin.Context, req *validatingOptimizer) (string, *RtnStatus) {
		return "ok", nil
	}
	h := bindAndRespond(func(c *gin.Context, req *validatingOptimizer) error {
		return c.ShouldBindJSON(req)
	}, handler)

	// 校验失败：返回第一条消息，Optimize 不被调用
	w := httptest.NewRecorder()
	ginCtx, _ := gin.CreateTestContext(w)
	ginCtx.Request = httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{}`))
	h(ginCtx)

	var resp struct {
		Code CodeType `json:"code"`
		Msg  string   `json:"msg"`
		Data any      `json:"data"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("unmarshal resp: %v", err)
	}
	if resp.Msg != "值不能为空" {
		t.Fatalf("expected first validation message, got %q (body=%s)", resp.Msg, w.Body.String())
	}
	if resp.Code != CodeError {
		t.Fatalf("expected CodeError, got %d", resp.Code)
	}

	// 校验通过：执行 Optimize 与 handler
	w2 := httptest.NewRecorder()
	ginCtx2, _ := gin.CreateTestContext(w2)
	ginCtx2.Request = httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"value":"x"}`))
	h(ginCtx2)

	var okResp struct {
		Code CodeType `json:"code"`
		Msg  string   `json:"msg"`
	}
	if err := json.Unmarshal(w2.Body.Bytes(), &okResp); err != nil {
		t.Fatalf("unmarshal ok resp: %v", err)
	}
	if okResp.Code != CodeSuccess {
		t.Fatalf("expected success, got %#v", okResp)
	}
}

// TestBindAndRespondMapsParseErrorSafely 断言：JSON 语法错误走安全映射返回“参数错误”。
func TestBindAndRespondMapsParseErrorSafely(t *testing.T) {
	gin.SetMode(gin.TestMode)

	handler := func(c *gin.Context, req *validatingOptimizer) (string, *RtnStatus) {
		return "ok", nil
	}
	h := bindAndRespond(func(c *gin.Context, req *validatingOptimizer) error {
		return c.ShouldBindJSON(req)
	}, handler)

	w := httptest.NewRecorder()
	ginCtx, _ := gin.CreateTestContext(w)
	ginCtx.Request = httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{not-json`))
	h(ginCtx)

	var resp struct {
		Code CodeType `json:"code"`
		Msg  string   `json:"msg"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("unmarshal resp: %v", err)
	}
	if resp.Msg != "参数错误" {
		t.Fatalf("expected safe 参数错误 mapping, got %q", resp.Msg)
	}
}

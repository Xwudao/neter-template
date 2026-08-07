package core

import (
	"errors"
	"testing"

	"github.com/Xwudao/neter-template/internal/domain/errs"
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

// messageValidator 提供字段级自定义消息。
type messageValidator struct {
	Value string `binding:"required"`
}

func (m *messageValidator) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"Value.required": "值不能为空",
	}
}

func TestDefaultBindErrorMapperUsesCustomMessage(t *testing.T) {
	err := defaultBindErrorMapper(&messageValidator{}, errors.New("unused"))
	// 非 ValidationErrors 走兜底"参数错误"
	if err == nil || err.Error() != "参数错误" {
		t.Fatalf("unexpected err: %v", err)
	}
}

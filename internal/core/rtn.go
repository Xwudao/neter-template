package core

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/Xwudao/neter-template/internal/data/ent"
	"github.com/Xwudao/neter-template/internal/domain/errs"
	"github.com/Xwudao/neter-template/internal/validate"
)

type CodeType int

const (
	// CodeSuccess 请求成功
	CodeSuccess CodeType = 200
	// CodeError 请求失败
	CodeError CodeType = 0
	// CodePassError 密码错误
	CodePassError CodeType = 400
	// CodeInvalid 网盘资源失效
	CodeInvalid CodeType = 401
	// CodeUnAuth 未授权
	CodeUnAuth CodeType = 403
	// CodeLogin 需要登录
	CodeLogin CodeType = 405
	// CodeWechatPwd 微信密码错误
	CodeWechatPwd CodeType = 406
	// CodeSignError 签名错误
	CodeSignError CodeType = 407
	// CodeCoinNotEnough 积分不足
	CodeCoinNotEnough CodeType = 408
	// CodeNeedVIP 需要VIP
	CodeNeedVIP CodeType = 409
	// NeedBuy 需要购买
	NeedBuy CodeType = 410
)

type RtnStatus struct {
	Code    CodeType `json:"code"`
	Message string   `json:"message"`
}

func NewRtnStatus(code CodeType, message string) *RtnStatus {
	return &RtnStatus{Code: code, Message: message}
}

func NewRtnWithErr(err error) *RtnStatus {
	var msg = err.Error()
	var businessErr *errs.Error

	switch {
	case errors.As(err, &businessErr):
		msg = businessErr.Message
	case ent.IsNotFound(err):
		msg = "记录不存在"
	case strings.Contains(msg, "Duplicate entry"):
		return &RtnStatus{
			Code:    CodeError,
			Message: "资源已存在",
		}
	case strings.Contains(msg, "connection refused"):
		return &RtnStatus{
			Code:    CodeError,
			Message: "内部网络连接错误",
		}
	}

	return &RtnStatus{
		Code:    0,
		Message: msg,
	}
}

var (
	Success = &RtnStatus{200, "请求成功"}
	Fail    = &RtnStatus{0, "请求失败"}
)

// WrappedResp 被包装的输出结构
type WrappedResp struct {
	Code CodeType `json:"code"`
	Msg  string   `json:"msg"`
	Data any      `json:"data"`
}

// Handler declares an API contract directly at the handler boundary.  Gin
// receives an ordinary HandlerFunc, while Request and Response remain visible
// to generators and static tooling.
type Handler[Request any, Response any] func(*gin.Context, *Request) (Response, *RtnStatus)

// ErrHandler 与 Handler 等价，但第二个返回值是 error：
// 返回 error 时代码自动转为 RtnStatus（可用 RtnStatusError 携带业务码）。
// 相比 Handler，业务函数体不再需要写 `return x, core.NewRtnWithErr(err)`。
type ErrHandler[Request any, Response any] func(*gin.Context, *Request) (Response, error)

// EmptyResponse 无数据成功响应，序列化为 `data: null`。
type EmptyResponse struct{}

// ListResponse 列表响应的标准结构（list/total）。
type ListResponse[T any] struct {
	List  []T   `json:"list"`
	Total int64 `json:"total"`
}

// BindErrorMapper 将请求绑定错误映射为用户可读的错误信息。
type BindErrorMapper func(request any, err error) error

// RequestValidator 由请求结构体实现，提供代码式字段校验。
// 框架在绑定成功、Optimize 之前自动调用；校验失败返回错误，不执行 handler。
type RequestValidator interface {
	Validate() error
}

// defaultBindErrorMapper 默认绑定错误映射：
// 绑定解析失败（JSON 语法/类型错误、空 body、strconv 错误）统一返回"参数错误"，
// 不把原始解析错误暴露给客户端；字段校验失败由 Validate() 在绑定成功后处理。
func defaultBindErrorMapper(request any, err error) error {
	return errors.New("参数错误")
}

// Optimizer 绑定成功后自动调用的接口（无返回值）。
// 请求结构体实现后，框架在绑定完成、handler 执行前自动调用，无需在 handler 内手动调用。
type Optimizer interface {
	Optimize()
}

// OptimizerErr 绑定成功后自动调用的接口（返回 error）。
// 优化失败时框架直接返回错误响应，不执行 handler。
type OptimizerErr interface {
	Optimize() error
}

// RtnStatusError 携带业务码的错误，用于 ErrHandler 中返回特殊业务码。
type RtnStatusError struct {
	Status *RtnStatus
}

func (e *RtnStatusError) Error() string {
	if e.Status == nil {
		return ""
	}
	return e.Status.Message
}

// NewRtnError 构造携带业务码的错误（供 ErrHandler 返回特殊业务码，如 CodeNeedVIP）。
func NewRtnError(code CodeType, msg string) error {
	return &RtnStatusError{Status: NewRtnStatus(code, msg)}
}

// errToRtn 将 error 转为 RtnStatus；识别 RtnStatusError 时直接使用其业务码。
func errToRtn(err error) *RtnStatus {
	var rtnErr *RtnStatusError
	if errors.As(err, &rtnErr) && rtnErr.Status != nil {
		return rtnErr.Status
	}
	return NewRtnWithErr(err)
}

// errHandlerToHandler 将 ErrHandler 适配为 Handler（内部统一走 errToRtn）。
func errHandlerToHandler[Request, Response any](h ErrHandler[Request, Response]) Handler[Request, Response] {
	return func(c *gin.Context, req *Request) (Response, *RtnStatus) {
		resp, err := h(c, req)
		if err != nil {
			return resp, errToRtn(err)
		}
		return resp, nil
	}
}

func JSON[Request any, Response any](handler Handler[Request, Response], mappers ...BindErrorMapper) gin.HandlerFunc {
	return bindAndRespond(func(c *gin.Context, request *Request) error {
		return c.ShouldBindJSON(request)
	}, handler, mappers...)
}

func Request[Request any, Response any](handler Handler[Request, Response], mappers ...BindErrorMapper) gin.HandlerFunc {
	return bindAndRespond(func(c *gin.Context, request *Request) error {
		return c.ShouldBind(request)
	}, handler, mappers...)
}

// JSONE 绑定 JSON body 后调用 ErrHandler；绑定错误与业务 error 自动转响应。
func JSONE[Req any, Resp any](handler ErrHandler[Req, Resp], mappers ...BindErrorMapper) gin.HandlerFunc {
	return JSON(errHandlerToHandler[Req, Resp](handler), mappers...)
}

// RequestE 绑定 query/form 后调用 ErrHandler；绑定错误与业务 error 自动转响应。
func RequestE[Req any, Resp any](handler ErrHandler[Req, Resp], mappers ...BindErrorMapper) gin.HandlerFunc {
	return Request(errHandlerToHandler[Req, Resp](handler), mappers...)
}

func NoInput[Response any](handler func(*gin.Context) (Response, *RtnStatus)) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, status := handler(c)
		writeResponse(c, data, status)
	}
}

// NoInputE 无绑定调用 ErrHandler；业务 error 自动转响应。
func NoInputE[Response any](handler func(*gin.Context) (Response, error)) gin.HandlerFunc {
	return NoInput(func(c *gin.Context) (Response, *RtnStatus) {
		resp, err := handler(c)
		if err != nil {
			return resp, errToRtn(err)
		}
		return resp, nil
	})
}

func bindAndRespond[Request any, Response any](bind func(*gin.Context, *Request) error, handler Handler[Request, Response], mappers ...BindErrorMapper) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request Request
		if err := bind(c, &request); err != nil {
			err = applyBindMappers(&request, err, mappers)
			var zero Response
			writeResponse(c, zero, NewRtnWithErr(err))
			return
		}
		if err := validateRequest(&request); err != nil {
			var zero Response
			writeResponse(c, zero, NewRtnWithErr(err))
			return
		}
		if !optimizeRequest(&request) {
			var zero Response
			writeResponse(c, zero, NewRtnWithErr(errors.New("参数优化失败")))
			return
		}
		data, status := handler(c, &request)
		writeResponse(c, data, status)
	}
}

// applyBindMappers 应用绑定错误映射；未显式传 mapper 时使用默认映射。
func applyBindMappers(request any, err error, mappers []BindErrorMapper) error {
	if len(mappers) == 0 {
		return defaultBindErrorMapper(request, err)
	}
	for _, mapper := range mappers {
		if mapper != nil {
			err = mapper(request, err)
		}
	}
	return err
}

// validateRequest 绑定成功后自动调用请求结构体的 Validate 方法（若实现）。
// 校验失败时只向响应暴露第一条消息，维持既有 msg 行为。
func validateRequest[T any](request *T) error {
	if v, ok := any(request).(RequestValidator); ok {
		if err := v.Validate(); err != nil {
			return validate.FirstError(err)
		}
	}
	return nil
}

// optimizeRequest 绑定成功后自动调用请求结构体的 Optimize 方法（若实现）。
// 支持 `Optimize()` 与 `Optimize() error` 两种签名；返回 false 表示优化失败。
func optimizeRequest[T any](request *T) bool {
	if opt, ok := any(request).(OptimizerErr); ok {
		if err := opt.Optimize(); err != nil {
			return false
		}
		return true
	}
	if opt, ok := any(request).(Optimizer); ok {
		opt.Optimize()
	}
	return true
}

func writeResponse(c *gin.Context, data any, status *RtnStatus) {
	resp := new(WrappedResp)
	if status != nil {
		resp.Code = status.Code
		resp.Msg = status.Message
	} else {
		resp.Code = Success.Code
		resp.Msg = Success.Message
	}
	resp.Data = data
	c.JSON(http.StatusOK, resp)
}

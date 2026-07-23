package core

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/Xwudao/neter-template/internal/data/ent"
	"github.com/Xwudao/neter-template/internal/domain/errs"
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

func NewListRtn[T ~int | ~int64](data any, total T) (gin.H, *RtnStatus) {
	return gin.H{
			"list":  data,
			"total": total,
		}, &RtnStatus{
			Code:    CodeSuccess,
			Message: "ok",
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

type WrappedHandlerFunc func(*gin.Context) (any, *RtnStatus)

// Handler declares an API contract directly at the handler boundary.  Gin
// receives an ordinary HandlerFunc, while Request and Response remain visible
// to generators and static tooling.
type Handler[Request any, Response any] func(*gin.Context, *Request) (Response, *RtnStatus)

type EmptyResponse struct{}

func JSON[Request any, Response any](handler Handler[Request, Response]) gin.HandlerFunc {
	return bindAndRespond(func(c *gin.Context, request *Request) error {
		return c.ShouldBindJSON(request)
	}, handler)
}

func Request[Request any, Response any](handler Handler[Request, Response]) gin.HandlerFunc {
	return bindAndRespond(func(c *gin.Context, request *Request) error {
		return c.ShouldBind(request)
	}, handler)
}

func NoInput[Response any](handler func(*gin.Context) (Response, *RtnStatus)) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, status := handler(c)
		writeResponse(c, data, status)
	}
}

func bindAndRespond[Request any, Response any](bind func(*gin.Context, *Request) error, handler Handler[Request, Response]) gin.HandlerFunc {
	return func(c *gin.Context) {
		var request Request
		if err := bind(c, &request); err != nil {
			var zero Response
			writeResponse(c, zero, NewRtnWithErr(err))
			return
		}
		data, status := handler(c, &request)
		writeResponse(c, data, status)
	}
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

// WrapData 包装响应结果
func WrapData(handler WrappedHandlerFunc) func(*gin.Context) {
	return func(c *gin.Context) {
		data, stat := handler(c)
		writeResponse(c, data, stat)
	}
}

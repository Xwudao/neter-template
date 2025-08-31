package core

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/Xwudao/neter-template/internal/data/ent"
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

	switch {
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

// WrapData 包装响应结果
func WrapData(handler WrappedHandlerFunc) func(*gin.Context) {
	return func(c *gin.Context) {
		data, stat := handler(c)

		resp := new(WrappedResp)
		if stat != nil {
			resp.Code = stat.Code
			resp.Msg = stat.Message
		} else {
			resp.Code = Success.Code
			resp.Msg = Success.Message
		}
		resp.Data = data
		c.JSON(http.StatusOK, resp)
	}
}

package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RtnStatus struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewRtnStatus(code int, message string) *RtnStatus {
	return &RtnStatus{Code: code, Message: message}
}

var (
	Success = &RtnStatus{200, "请求成功"}
	Fail    = &RtnStatus{0, "请求失败"}
)

// WrappedResp 被包装的输出结构
type WrappedResp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type WrappedHandlerFunc func(*gin.Context) (interface{}, *RtnStatus)

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

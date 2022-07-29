package mdw

import (
	"github.com/gin-gonic/gin"
	"time"
)

type RouterLogFields struct {
	Status int           `json:"status"`
	Time   time.Duration `json:"time"`
	Method string        `json:"method"`
	Path   string        `json:"path"`
	IP     string        `json:"ip"`
	Agent  string        `json:"agent"`
	Uri    string        `json:"uri"`
}

func LoggerMiddleware(apply func(fields *RouterLogFields)) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		start := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		end := time.Now()
		//执行时间
		latency := end.Sub(start)

		path := c.Request.URL.Path
		uri := c.Request.RequestURI

		clientIP := c.GetHeader("X-Forwarded-For")
		if clientIP == "" {
			clientIP = c.GetHeader("X-Real-IP")
		}
		if clientIP == "" {
			clientIP = c.ClientIP()
		}
		method := c.Request.Method
		statusCode := c.Writer.Status()

		fields := &RouterLogFields{
			Status: statusCode,
			Time:   latency,
			Method: method,
			Path:   path,
			IP:     clientIP,
			Agent:  c.Request.UserAgent(),
			Uri:    uri,
		}

		apply(fields)
	}
}

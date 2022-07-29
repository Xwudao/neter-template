package mdw

import (
	"github.com/gin-gonic/gin"
	"net"
	"strings"
)

func IPLimitMiddleware(countPerSeconds int, limited func(c *gin.Context, ip string)) gin.HandlerFunc {
	limiter := NewIPRateLimiter(countPerSeconds)
	return func(c *gin.Context) {
		ipAddr := getIP(c)
		lt := limiter.GetLimiter(ipAddr)
		if !lt.Allow() {
			limited(c, ipAddr)
			c.Abort()
			return
		} else {
			c.Next()
		}
	}
}

func getIP(c *gin.Context) string {
	ip, _, err := net.SplitHostPort(c.Request.RemoteAddr)
	if err != nil {
		ip = c.Request.RemoteAddr
	}
	if ip != "127.0.0.1" {
		return ip
	}
	// Check if behide nginx or apache
	xRealIP := c.Request.Header.Get("X-Real-Ip")
	xForwardedFor := c.Request.Header.Get("X-Forwarded-For")

	for _, address := range strings.Split(xForwardedFor, ",") {
		address = strings.TrimSpace(address)
		if address != "" {
			return address
		}
	}

	if xRealIP != "" {
		return xRealIP
	}
	return ip
}

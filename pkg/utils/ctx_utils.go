package utils

import (
	"net"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/Xwudao/neter-template/internal/data/sqlc"
	"github.com/Xwudao/neter-template/pkg/enum"
)

func GetIP(c *gin.Context) string {
	// Check if behind nginx or apache
	xRealIP := c.Request.Header.Get("X-Real-Ip")
	xForwardedFor := c.Request.Header.Get("X-Forwarded-For")

	if xRealIP != "" {
		return xRealIP
	}

	// Split X-Forwarded-For header and consider the first IP as the client's IP
	clientIP := handleFirstIp(xForwardedFor)

	// If no IP found in headers, use RemoteAddr
	if clientIP == "" {
		clientIP, _, _ = net.SplitHostPort(c.Request.RemoteAddr)
	}

	// Validate the extracted IP address format
	parsedIP := net.ParseIP(clientIP)
	if parsedIP == nil {
		return "" // or handle the error accordingly
	}

	return clientIP
}

func handleFirstIp(ips string) string {
	if ips == "" {
		return ""
	}
	ipsSlice := strings.Split(ips, ",")
	return strings.TrimSpace(ipsSlice[0])
}

func IsAdmin(c *gin.Context) bool {
	u, ok := c.Get(enum.KeyUserInfo)
	if !ok {
		return false
	}

	userInfo := u.(*sqlc.User)

	return userInfo.Role == sqlc.UserRoleAdmin
}

// User 获取用户信息，可能为nil
func User(c *gin.Context) *sqlc.User {
	vl, exists := c.Get(enum.KeyUserInfo)
	if !exists {
		return nil
	}

	return vl.(*sqlc.User)
}

func UserID(c *gin.Context) int64 {
	u, ok := c.Get(enum.KeyUserInfo)
	if !ok {
		return 0
	}

	userInfo := u.(*sqlc.User)

	return userInfo.ID
}

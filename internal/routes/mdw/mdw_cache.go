package mdw

import (
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func CacheMdw() gin.HandlerFunc {
	return func(c *gin.Context) {
		ext := filepath.Ext(c.Request.URL.Path)
		if ext == "" {
			return
		}

		switch strings.ToLower(ext) {
		case ".js", ".css", ".png", ".jpg", ".jpeg":
			c.Header("Cache-Control", "public, max-age=86400")
		}
	}
}

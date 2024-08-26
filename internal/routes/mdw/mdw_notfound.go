package mdw

import (
	"mime"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func NotFoundMdw() gin.HandlerFunc {
	return func(c *gin.Context) {
		var path = c.Request.URL.Path
		if hasStaticFile(path) {
			cnt, mimeType := readStaticFile(path)
			c.Data(200, mimeType, cnt)
			c.Abort()
			return
		}

		c.String(404, "404 page not found")
	}
}

func hasStaticFile(path string) bool {
	pw, _ := os.Getwd()
	fp := filepath.Join(pw, "web/public", path)
	info, err := os.Stat(fp)
	if err != nil {
		return false
	}

	return !info.IsDir()
}

func readStaticFile(path string) ([]byte, string) {
	pw, _ := os.Getwd()
	fp := filepath.Join(pw, "web/public", path)
	cnt, err := os.ReadFile(fp)
	if err != nil {
		return nil, ""
	}
	mimeType := mime.TypeByExtension(filepath.Ext(fp))

	return cnt, mimeType
}

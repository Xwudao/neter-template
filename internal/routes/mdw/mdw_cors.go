package mdw

import (
	"path"
	"path/filepath"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/knadh/koanf/v2"

	"github.com/Xwudao/neter-template/pkg/enum"
)

func CorsMdw(conf *koanf.Koanf) gin.HandlerFunc {
	//config cors
	allowHeaders := enum.AllowHeaders
	exposeHeaders := enum.ExposeHeaders
	allowCredentials := conf.Bool("cors.allowCredentials")
	allowOrigins := conf.Strings("cors.allowOrigin")
	allowAllOrigins := len(allowOrigins) == 0 && !allowCredentials

	methods := []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	maxAge := conf.Duration("cors.maxAge")

	var cs = cors.Config{
		//AllowAllOrigins:  false,
		//AllowOrigins:     allowOrigins,
		AllowMethods:     methods,
		AllowWildcard:    true,
		AllowHeaders:     allowHeaders,
		AllowCredentials: allowCredentials,
		ExposeHeaders:    exposeHeaders,
		MaxAge:           maxAge,
		// 用了 AllowOriginFunc 就不能设置 AllowAllOrigins 和 AllowOrigins 了
		AllowOriginWithContextFunc: func(c *gin.Context, origin string) bool {
			if allowAllOrigins || isStaticCORSPath(c.Request.URL.Path) {
				return true
			}
			for _, allowOrigin := range allowOrigins {
				if ok, _ := path.Match(allowOrigin, origin); ok {
					return true
				}
			}

			return false
		},
	}

	return cors.New(cs)
}

func isStaticCORSPath(requestPath string) bool {
	if strings.HasPrefix(requestPath, "/assets/") || strings.HasPrefix(requestPath, "/static/") {
		return true
	}

	switch strings.ToLower(filepath.Ext(requestPath)) {
	case ".css", ".js", ".map", ".png", ".jpg", ".jpeg", ".gif", ".svg", ".webp", ".ico", ".woff", ".woff2", ".ttf":
		return true
	default:
		return false
	}
}

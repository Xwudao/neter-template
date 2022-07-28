package v1

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/knadh/koanf"
)

type HomeRoute struct {
	router *gin.Engine
	conf   *koanf.Koanf
}

func NewHomeRoute(router *gin.Engine, conf *koanf.Koanf) *HomeRoute {
	r := &HomeRoute{router: router, conf: conf}

	r.router.GET("/", r.home())

	return r
}

func (r *HomeRoute) home() gin.HandlerFunc {
	return func(c *gin.Context) {
		port := r.conf.Int("port")

		c.String(200, "Hello World!"+strconv.Itoa(port))
	}
}

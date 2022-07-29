package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/knadh/koanf"

	"github.com/Xwudao/neter-template/internal/biz"
)

type HomeRoute struct {
	conf *koanf.Koanf
	hb   *biz.HomeBiz
	g    *gin.Engine
}

func NewHomeRoute(g *gin.Engine, conf *koanf.Koanf, hb *biz.HomeBiz) *HomeRoute {
	r := &HomeRoute{
		conf: conf,
		hb:   hb,
		g:    g,
	}

	return r
}

func (r *HomeRoute) Reg() {
	r.g.GET("/", r.home())
}

func (r *HomeRoute) home() gin.HandlerFunc {
	return func(c *gin.Context) {
		name := c.Query("name")

		c.String(200, r.hb.SayHello(name))
	}
}

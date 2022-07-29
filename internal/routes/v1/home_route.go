package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/knadh/koanf"

	"github.com/Xwudao/neter-template/internal/biz"
	"github.com/Xwudao/neter-template/internal/core"
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
	r.g.GET("/v1", core.WrapData(r.home()))
}

func (r *HomeRoute) home() core.WrappedHandlerFunc {
	return func(c *gin.Context) (interface{}, *core.RtnStatus) {
		return "hello", nil
	}
}

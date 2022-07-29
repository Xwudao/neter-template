package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/knadh/koanf"

	"github.com/Xwudao/neter-template/internal/biz"
	"github.com/Xwudao/neter-template/internal/core"
	"github.com/Xwudao/neter-template/pkg/utils/jwt"
)

type HomeRoute struct {
	conf *koanf.Koanf
	hb   *biz.HomeBiz
	g    *gin.Engine
	jwt  *jwt.Client
}

func NewHomeRoute(g *gin.Engine, jwt *jwt.Client, conf *koanf.Koanf, hb *biz.HomeBiz) *HomeRoute {
	r := &HomeRoute{
		conf: conf,
		hb:   hb,
		g:    g,
		jwt:  jwt,
	}

	return r
}

func (r *HomeRoute) Reg() {
	r.g.GET("/v1", core.WrapData(r.home()))
}

func (r *HomeRoute) home() core.WrappedHandlerFunc {
	return func(c *gin.Context) (interface{}, *core.RtnStatus) {
		tokenString, err := r.jwt.Generate(1)
		if err != nil {
			return nil, core.Fail
		}
		return tokenString, nil
	}
}

package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/knadh/koanf/v2"

	"github.com/Xwudao/neter-template/internal/biz"
	"github.com/Xwudao/neter-template/internal/core"
	"github.com/Xwudao/neter-template/internal/routes/mdw"
)

type UserRoute struct {
	conf *koanf.Koanf
	g    *gin.Engine
	uz   *biz.UserBiz
}

func NewUserRoute(g *gin.Engine, uz *biz.UserBiz, conf *koanf.Koanf) *UserRoute {
	r := &UserRoute{
		conf: conf,
		g:    g,
		uz:   uz,
	}

	return r
}

func (r *UserRoute) Reg() {
	// r.g.GET("/v1/user", core.WrapData(r.user()))

	group := r.g.Group("/v1/user")
	{
		group.GET("", core.WrapData(r.user()))
	}
	authGroup := r.g.Group("/auth/v1/user").Use(mdw.MustLoginMiddleware())
	{
		// authGroup.GET("/auth", core.WrapData(r.user()))
		_ = authGroup
	}
	adminGroup := r.g.Group("/admin/v1/user").Use(mdw.MustWithRoleMiddleware("admin"))
	{
		_ = adminGroup
	}
}

func (r *UserRoute) user() core.WrappedHandlerFunc {
	return func(c *gin.Context) (any, *core.RtnStatus) {
		//var (
		//	ctx = c.Request.Context()
		//)
		return "hello", nil
	}
}

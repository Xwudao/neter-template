package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/knadh/koanf/v2"

	"github.com/Xwudao/neter-template/internal/biz"
	"github.com/Xwudao/neter-template/internal/core"
	"github.com/Xwudao/neter-template/internal/data/ent/user"
	"github.com/Xwudao/neter-template/internal/domain/params"
	"github.com/Xwudao/neter-template/internal/routes/mdw"
	"github.com/Xwudao/neter-template/internal/routes/valid"
	"github.com/Xwudao/neter-template/pkg/utils"
)

type UserRoute struct {
	conf *koanf.Koanf
	g    *gin.Engine
	ub   *biz.UserBiz
}

func NewUserRoute(g *gin.Engine, uz *biz.UserBiz, conf *koanf.Koanf) *UserRoute {
	r := &UserRoute{
		conf: conf,
		g:    g,
		ub:   uz,
	}

	return r
}

func (r *UserRoute) Reg() {
	// r.g.GET("/v1/user", core.WrapData(r.user()))

	group := r.g.Group("/v1/user")
	{
		group.GET("", core.WrapData(r.user()))
		group.POST("/login", core.WrapData(r.login()))
	}
	authGroup := r.g.Group("/auth/v1/user").Use(mdw.MustLoginMiddleware())
	{
		authGroup.GET("/info", core.WrapData(r.info()))
	}
	adminGroup := r.g.Group("/admin/v1/user").Use(mdw.MustWithRoleMiddleware(user.RoleAdmin))
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

func (r *UserRoute) login() core.WrappedHandlerFunc {
	return func(c *gin.Context) (any, *core.RtnStatus) {
		var (
			ctx = c.Request.Context()
			pm  params.UserLoginParams
		)
		if err := c.ShouldBindJSON(&pm); err != nil {
			return nil, core.NewRtnWithErr(valid.GetErrorMsg(&pm, err))
		}

		u, token, err := r.ub.Login(ctx, &pm)
		if err != nil {
			return nil, core.NewRtnWithErr(err)
		}

		return gin.H{
			"user":  u,
			"token": token,
		}, nil
	}
}

func (r *UserRoute) info() core.WrappedHandlerFunc {
	return func(c *gin.Context) (any, *core.RtnStatus) {
		us := utils.User(c)

		return us, nil
	}
}

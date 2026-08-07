package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/knadh/koanf/v2"

	"github.com/Xwudao/neter-template/internal/biz"
	"github.com/Xwudao/neter-template/internal/core"
	"github.com/Xwudao/neter-template/internal/data/ent"
	"github.com/Xwudao/neter-template/internal/data/ent/user"
	"github.com/Xwudao/neter-template/internal/domain/params"
	"github.com/Xwudao/neter-template/internal/routes/mdw"
	"github.com/Xwudao/neter-template/pkg/utils"
)

type UserRoute struct {
	conf *koanf.Koanf
	ub   biz.UserBizIface
}

func NewUserRoute(uz biz.UserBizIface, conf *koanf.Koanf) *UserRoute {
	r := &UserRoute{
		conf: conf,
		ub:   uz,
	}

	return r
}

type UserLoginResponse struct {
	User  *ent.User `json:"user"`
	Token string    `json:"token"`
}

func (r *UserRoute) Register(router gin.IRouter) {
	// router.GET("/v1/user", core.NoInputE(r.user))

	group := router.Group("/v1/user")
	{
		group.GET("", core.NoInputE(r.user))
		group.POST("/login", core.JSONE(r.login))
	}
	authGroup := router.Group("/auth/v1/user").Use(mdw.MustLoginMiddleware())
	{
		authGroup.GET("/info", core.NoInputE(r.info))
	}
	adminGroup := router.Group("/admin/v1/user").Use(mdw.MustWithRoleMiddleware(user.RoleAdmin))
	{
		_ = adminGroup
	}
}

func (r *UserRoute) user(c *gin.Context) (string, error) {
	return "hello", nil
}

func (r *UserRoute) login(c *gin.Context, pm *params.UserLoginParams) (UserLoginResponse, error) {
	u, token, err := r.ub.Login(c.Request.Context(), pm)
	if err != nil {
		return UserLoginResponse{}, err
	}
	return UserLoginResponse{User: u, Token: token}, nil
}

func (r *UserRoute) info(c *gin.Context) (*ent.User, error) {
	return utils.User(c), nil
}

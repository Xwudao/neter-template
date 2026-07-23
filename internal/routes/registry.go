package routes

import (
	"github.com/gin-gonic/gin"

	v1 "github.com/Xwudao/neter-template/internal/routes/v1"
)

// Registrar is the only lifecycle contract HttpEngine needs from a route.
// Individual routes still own their Gin groups and middleware.
type Registrar interface {
	Register(gin.IRouter)
}

type RouteRegistry []Registrar

func NewRouteRegistry(
	userRoute *v1.UserRoute,
	siteConfigRoute *v1.SiteConfigRoute,
	dataListRoute *v1.DataListRoute,
) RouteRegistry {
	return RouteRegistry{
		userRoute,
		siteConfigRoute,
		dataListRoute,
	}
}

func (r RouteRegistry) RegisterAll(router gin.IRouter) {
	for _, route := range r {
		route.Register(router)
	}
}

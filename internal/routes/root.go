package routes

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	v1 "github.com/Xwudao/neter/internal/routes/v1"
)

func NewEngine() *gin.Engine {
	r := gin.Default()

	return r
}

type AppRoutes struct {
	log    *zap.SugaredLogger
	v1Home *v1.HomeRoute
}

func NewAppRoutes(log *zap.SugaredLogger, v1Home *v1.HomeRoute) *AppRoutes {
	return &AppRoutes{v1Home: v1Home, log: log}
}

func (r *AppRoutes) Setup() {
	r.log.Infof("app routes setup")
}

package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/knadh/koanf"
	"go.uber.org/zap"

	v1 "github.com/Xwudao/neter-template/internal/routes/v1"
)

func NewEngine(conf *koanf.Koanf) *gin.Engine {

	mode := conf.String("app.mode")
	if mode != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	return r
}

type HttpEngine struct {
	router *gin.Engine
	conf   *koanf.Koanf
	log    *zap.SugaredLogger

	homeRoute *v1.HomeRoute
}

func NewHttpEngine(
	router *gin.Engine,
	conf *koanf.Koanf,
	log *zap.SugaredLogger,
	homeRoute *v1.HomeRoute,
) (*HttpEngine, error) {

	he := &HttpEngine{
		conf:   conf,
		log:    log,
		router: router,

		homeRoute: homeRoute,
	}

	return he, nil
}

func (r *HttpEngine) Run() error {
	conf := r.conf
	log := r.log
	router := r.router

	port := conf.Int("app.port")
	err := router.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	log.Infof("app running on port: %d", port)

	return nil
}
func (r *HttpEngine) Register() {
	r.homeRoute.Reg()
}

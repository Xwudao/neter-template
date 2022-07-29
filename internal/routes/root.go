package routes

import (
	"fmt"

	"github.com/Xwudao/neter-template/internal/routes/mdw"
	v1 "github.com/Xwudao/neter-template/internal/routes/v1"

	"github.com/gin-gonic/gin"
	"github.com/knadh/koanf"
	"go.uber.org/zap"
)

func NewEngine(conf *koanf.Koanf, log *zap.SugaredLogger) *gin.Engine {

	mode := conf.String("app.mode")
	if mode != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}

	logFunc := func(fields *mdw.RouterLogFields) {
		log.Infow("visit",
			zap.Int("status", fields.Status),
			zap.String("method", fields.Method),
			zap.String("time", fields.Time.String()),
			zap.String("path", fields.Path),
			zap.String("ip", fields.IP),
			zap.String("agent", fields.Agent),
			zap.String("uri", fields.Uri))
	}

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery(), mdw.LoggerMiddleware(logFunc))

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

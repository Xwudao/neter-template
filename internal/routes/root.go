package routes

import (
	"fmt"
	v1 "github.com/Xwudao/neter-template/internal/routes/v1"
	"github.com/gin-gonic/gin"
	"github.com/knadh/koanf"
	"go.uber.org/zap"
)

type HttpEngine struct {
	router *gin.Engine
	conf   *koanf.Koanf
	log    *zap.SugaredLogger
}

func NewHttpEngine(conf *koanf.Koanf, log *zap.SugaredLogger) *HttpEngine {

	he := &HttpEngine{
		conf: conf,
		log:  log,
	}

	return he
}

func (h *HttpEngine) Run() error {
	mode := h.conf.String("app.mode")
	if mode != "debug" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	h.router = r
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	port := h.conf.Int("app.port")
	err := r.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	h.log.Infof("app running on port: %d", port)
	return nil
}

func (h *HttpEngine) Setup() {
	v1.NewHomeRoute(h.router, h.conf)
}

package routes

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Xwudao/neter-template/internal/core"
	"github.com/Xwudao/neter-template/internal/routes/mdw"
	v1 "github.com/Xwudao/neter-template/internal/routes/v1"
	"github.com/Xwudao/neter-template/pkg/config"
	"github.com/gin-contrib/cors"

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
	_ = r.SetTrustedProxies(nil)
	r.Use(gin.Logger())
	r.Use(gin.Recovery(), mdw.LoggerMiddleware(logFunc))

	return r
}

type HttpEngine struct {
	router *gin.Engine
	conf   *config.Config
	log    *zap.SugaredLogger
	ctx    *core.AppContext

	corsConf cors.Config

	v1UserRoute *v1.UserRoute
}

func NewHttpEngine(
	router *gin.Engine,
	conf *config.Config,
	log *zap.SugaredLogger,
	ctx *core.AppContext,
	v1UserRoute *v1.UserRoute,
) (*HttpEngine, error) {

	he := &HttpEngine{
		conf:     conf,
		log:      log,
		router:   router,
		ctx:      ctx,
		corsConf: cors.DefaultConfig(),

		v1UserRoute: v1UserRoute,
	}

	return he, nil
}

func (r *HttpEngine) Run() error {
	conf := r.conf
	log := r.log
	router := r.router

	port := conf.App.Port
	addr := fmt.Sprintf(":%d", port)
	log.Infof("app running on port: %d", port)

	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	go func() {

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx := r.ctx.Ctx
	cancel := r.ctx.Cancel

	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Infof("server exiting")
	return nil
}
func (r *HttpEngine) Register() {
	r.v1UserRoute.Reg()
}

func (r *HttpEngine) ConfigCors() {
	r.router.Use(cors.New(r.corsConf))
}

func (r *HttpEngine) SetOriginFun(f func(origin string) bool) {
	r.corsConf.AllowOriginFunc = f
}
func (r *HttpEngine) SetMaxAge(age time.Duration) {
	r.corsConf.MaxAge = age
}
func (r *HttpEngine) SetCredentials(b bool) {
	r.corsConf.AllowCredentials = b
}
func (r *HttpEngine) SetExposeHeaders(s []string) {
	r.corsConf.ExposeHeaders = s
}
func (r *HttpEngine) SetHeaders(s []string) {
	r.corsConf.AllowHeaders = s
}
func (r *HttpEngine) SetMethods(s []string) {
	r.corsConf.AllowMethods = s
}
func (r *HttpEngine) SetOrigin(s []string) {
	r.corsConf.AllowOrigins = s
}

package routes

import (
	"context"
	"errors"
	"fmt"
	"mime"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/knadh/koanf/v2"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/Xwudao/neter-template/assets"
	"github.com/Xwudao/neter-template/internal/biz"
	"github.com/Xwudao/neter-template/internal/routes/mdw"
	"github.com/Xwudao/neter-template/internal/system"
	"github.com/Xwudao/neter-template/pkg/logger"
	"github.com/Xwudao/neter-template/pkg/utils/jwt"
)

func NewEngine(
	zw *logger.ZapWriter,
	jwt *jwt.Client,
	ur biz.UserRepository,
	conf *koanf.Koanf,
	sb *biz.SeoBizBiz,
	log *zap.SugaredLogger,
	ssr *SSRRenderer,
) (*gin.Engine, error) {
	var (
		isDebug   = conf.String("app.mode") == "debug"
		isRelease = conf.String("app.mode") == "release"
	)
	if isRelease {
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

	_ = mime.AddExtensionType(".js", "application/javascript")

	r := gin.New()
	_ = r.SetTrustedProxies(nil)

	r.Use(mdw.CorsMdw(conf))

	r.Use(mdw.CacheMdw(), mdw.ExtractUserInfoMiddleware(log, jwt, ur))

	spa, err := mdw.NewSpaMdw(assets.SpaDist, "dist", sb)
	if err != nil {
		return nil, err
	}
	r.NoRoute(ssr.Serve(spa.Serve("/")))
	//r.NoRoute(spa.Serve("index.html"))
	//r.NoRoute(mdw.NotFoundMdw())

	r.Use(mdw.DumpReqResMdw(isDebug, log))
	r.Use(gin.Logger())
	r.Use(gin.RecoveryWithWriter(zw), mdw.LoggerMiddleware(logFunc))

	return r, nil
}

type HttpEngine struct {
	router *gin.Engine
	conf   *koanf.Koanf
	log    *zap.SugaredLogger
	ctx    *system.AppContext

	routes RouteRegistry
	ssr    *SSRRenderer
}

func NewHttpEngine(
	router *gin.Engine,
	conf *koanf.Koanf,
	log *zap.SugaredLogger,
	ctx *system.AppContext,
	routes RouteRegistry,
	ssr *SSRRenderer,
) (*HttpEngine, error) {

	he := &HttpEngine{
		conf:   conf,
		log:    log,
		router: router,
		ctx:    ctx,
		routes: routes,
		ssr:    ssr,
	}

	return he, nil
}

func (r *HttpEngine) Run() error {
	log := r.log
	router := r.router
	defer r.ctx.Cancel()
	defer func() {
		if err := r.ssr.Shutdown(context.Background()); err != nil {
			log.Warnw("shutdown SSR engine", "error", err)
		}
	}()

	port := r.conf.Int("app.port")
	host := r.conf.Bool("app.host")

	addr := fmt.Sprintf("127.0.0.1:%d", port)

	if host {
		addr = fmt.Sprintf(":%d", port)
	}

	log.Infof("app running on: http://127.0.0.1:%d", port)

	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	serverErr := make(chan error, 1)
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverErr <- err
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(quit)

	select {
	case err := <-serverErr:
		return fmt.Errorf("listen on %s: %w", addr, err)
	case <-quit:
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		return fmt.Errorf("shutdown server: %w", err)
	}
	log.Infof("server exiting")
	return nil
}
func (r *HttpEngine) Register() {
	r.routes.RegisterAll(r.router)
}

func (r *HttpEngine) Use(middleware ...gin.HandlerFunc) gin.IRoutes {
	return r.router.Use(middleware...)
}

func (r *HttpEngine) ConfigCors(c cors.Config) {
	r.router.Use(cors.New(c))
}

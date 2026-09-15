package routes

import (
	"context"
	"errors"
	"fmt"
	"mime"
	"net"
	"net/http"

	"github.com/Xwudao/loom"
	"github.com/knadh/koanf/v2"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/Xwudao/neter-template/assets"
	"github.com/Xwudao/neter-template/internal/biz"
	"github.com/Xwudao/neter-template/internal/cron"
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
	r.NoRoute(spa.Serve("/"))
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
	cron   *cron.Cron

	routes   RouteRegistry
	server   *http.Server
	serveErr chan error
}

func NewHttpEngine(
	router *gin.Engine,
	conf *koanf.Koanf,
	log *zap.SugaredLogger,
	ctx *system.AppContext,
	cron *cron.Cron,
	routes RouteRegistry,
	lifecycle *loom.Lifecycle,
) (*HttpEngine, error) {

	he := &HttpEngine{
		conf:   conf,
		log:    log,
		router: router,
		ctx:    ctx,
		cron:   cron,
		routes: routes,
	}

	lifecycle.Append(loom.Hook{OnStart: he.start, OnStop: he.stop})
	return he, nil
}

func (r *HttpEngine) start(context.Context) error {
	addr := fmt.Sprintf("127.0.0.1:%d", r.conf.Int("app.port"))
	if r.conf.Bool("app.host") {
		addr = fmt.Sprintf(":%d", r.conf.Int("app.port"))
	}
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("listen HTTP server on %s: %w", addr, err)
	}

	r.server = &http.Server{Addr: listener.Addr().String(), Handler: r.router}
	r.serveErr = make(chan error, 1)
	r.log.Infof("app running on: http://%s", listener.Addr())
	go func() {
		if err := r.server.Serve(listener); err != nil && !errors.Is(err, http.ErrServerClosed) {
			r.serveErr <- err
		}
	}()
	return nil
}

func (r *HttpEngine) stop(ctx context.Context) error {
	defer func() {
		r.ctx.Cancel()
		if r.cron != nil {
			_ = r.cron.Close()
		}
	}()
	if r.server == nil {
		return nil
	}
	if err := r.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("shutdown HTTP server: %w", err)
	}
	r.log.Infof("server exiting")
	return nil
}

func (r *HttpEngine) Wait(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return nil
	case err := <-r.serveErr:
		return err
	}
}

func (r *HttpEngine) CancelAppContext() { r.ctx.Cancel() }
func (r *HttpEngine) Register() {
	r.routes.RegisterAll(r.router)
}

func (r *HttpEngine) Use(middleware ...gin.HandlerFunc) gin.IRoutes {
	return r.router.Use(middleware...)
}

func (r *HttpEngine) ConfigCors(c cors.Config) {
	r.router.Use(cors.New(c))
}

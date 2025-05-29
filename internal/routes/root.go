package routes

import (
	"errors"
	"fmt"
	"mime"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-contrib/gzip"
	"github.com/knadh/koanf/v2"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"go-kitboxpro/assets"
	"go-kitboxpro/internal/biz"
	"go-kitboxpro/internal/data"
	"go-kitboxpro/internal/routes/mdw"
	v1 "go-kitboxpro/internal/routes/v1"
	"go-kitboxpro/internal/system"
	"go-kitboxpro/pkg/logger"
	"go-kitboxpro/pkg/utils/jwt"
)

func NewEngine(
	zw *logger.ZapWriter,
	jwt *jwt.Client,
	data *data.Data,
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

	r.Use(mdw.CacheMdw(), gzip.Gzip(gzip.DefaultCompression, gzip.WithExcludedPaths([]string{
		"/admin/",
		"/auth/",
		"/open/",
		"/v1/",
		"/v2/",
		"/v3/",
	})))

	r.Use(mdw.ExtractUserInfoMiddleware(log, jwt, data))

	//spa := mdw.NewSpaMdw(assets.SpaDist, "dist")
	//r.NoRoute(spa.ServeNotFound("index.html"))
	//r.NoRoute(mdw.NotFoundMdw())
	spa, err := mdw.NewSpaMdw(assets.SpaDist, "dist", sb)
	if err != nil {
		return nil, err
	}
	r.NoRoute(spa.Serve("/"))

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

	v1UserRoute       *v1.UserRoute
	v1SiteConfigRoute *v1.SiteConfigRoute
	v1DataListRoute   *v1.DataListRoute
}

func NewHttpEngine(
	router *gin.Engine,
	conf *koanf.Koanf,
	log *zap.SugaredLogger,
	ctx *system.AppContext,
	v1UserRoute *v1.UserRoute,
	v1SiteConfigRoute *v1.SiteConfigRoute,
	v1DataListRoute *v1.DataListRoute,
) (*HttpEngine, error) {

	he := &HttpEngine{
		conf:              conf,
		log:               log,
		router:            router,
		ctx:               ctx,
		v1UserRoute:       v1UserRoute,
		v1SiteConfigRoute: v1SiteConfigRoute,
		v1DataListRoute:   v1DataListRoute,
	}

	return he, nil
}

func (r *HttpEngine) Run() error {
	log := r.log
	router := r.router

	port := r.conf.Int("app.port")
	addr := fmt.Sprintf(":%d", port)

	log.Infof("app running on: http://127.0.0.1:%d", port)

	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
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
	r.v1SiteConfigRoute.Reg()
	r.v1DataListRoute.Reg()
}

func (r *HttpEngine) Use(middleware ...gin.HandlerFunc) gin.IRoutes {
	return r.router.Use(middleware...)
}

func (r *HttpEngine) ConfigCors(c cors.Config) {
	r.router.Use(cors.New(c))
}

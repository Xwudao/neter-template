package routes

import (
	"errors"
	"fmt"
	"mime"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/knadh/koanf/v2"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/Xwudao/neter-template/assets"
	"github.com/Xwudao/neter-template/internal/biz"
	"github.com/Xwudao/neter-template/internal/routes/mdw"
	v1 "github.com/Xwudao/neter-template/internal/routes/v1"
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

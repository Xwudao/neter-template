package cmd

import (
	"slices"

	"github.com/gin-contrib/cors"

	"github.com/Xwudao/neter-template/internal/biz"
	"github.com/Xwudao/neter-template/internal/cron"
	"github.com/Xwudao/neter-template/internal/system"
	"github.com/Xwudao/neter-template/pkg/enum"
	"github.com/Xwudao/neter-template/pkg/utils"
	"github.com/Xwudao/neter-template/pkg/varx"

	"github.com/knadh/koanf/v2"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/Xwudao/neter-template/internal/routes"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "",
	Short: "neter is a golang web framework",
	Long:  `neter is a golang web framework`,

	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute(run func(cmd *cobra.Command, args []string)) error {
	rootCmd.Run = run
	return rootCmd.Execute()
}

type MainApp struct {
	http       *routes.HttpEngine
	initSystem *system.InitSystem
	cron       *cron.Cron
	conf       *koanf.Koanf
	logger     *zap.SugaredLogger
	sib        *biz.SystemInitBiz
	scb        *biz.SiteConfigBiz
}

func NewMainApp(
	http *routes.HttpEngine,
	logger *zap.SugaredLogger,
	conf *koanf.Koanf,
	cron *cron.Cron,
	initSystem *system.InitSystem,
	sib *biz.SystemInitBiz,
	scb *biz.SiteConfigBiz,
) (*MainApp, func()) {
	m := &MainApp{
		logger:     logger,
		http:       http,
		initSystem: initSystem,
		cron:       cron,
		conf:       conf,
		sib:        sib, scb: scb,
	}

	cleanup := func() {
		logger.Infof("begin to cleanup")
		_ = m.cron.Close()
		logger.Infof("cleanup done")
	}

	m.checkSystem()

	return m, cleanup
}

func (m *MainApp) Run() error {
	m.initSystem.InitConfig()

	if err := m.sib.AddAdminUser(); err != nil {
		return err
	}
	if err := m.scb.Init(); err != nil {
		return err
	}

	m.cron.Run()
	m.cors() //this must be set before Register method
	m.http.Register()
	return m.http.Run()
}

// check system
func (m *MainApp) checkSystem() {
	if err := m.initSystem.CheckSystem(); err != nil {
		panic(err)
	}
}

func (m *MainApp) cors() {
	isDev := m.conf.String("app.mode") == "debug"

	credentials := m.conf.Bool("cors.allowCredentials")
	exposeHeaders := []string{"Content-Type", "Authorization", "X-Login"}
	methods := []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	maxAge := m.conf.Duration("cors.maxAge")
	domains := m.conf.Strings("cors.domains")

	var c = cors.Config{
		AllowOriginFunc: func(origin string) bool {
			if isDev {
				return varx.ContainLike([]string{"localhost", "127.0.0.1"}, origin)
			}
			domain, _ := utils.ExtractRootDomain(origin)
			//m.logger.Infof("cors check domain: %s, origin: %s", domain, origin)
			return slices.Contains(domains, domain)
		},
		AllowMethods:     methods,
		AllowHeaders:     enum.AllowHeaders,
		AllowCredentials: credentials,
		ExposeHeaders:    exposeHeaders,
		MaxAge:           maxAge,
	}

	m.http.ConfigCors(c)
	//config cors end
}

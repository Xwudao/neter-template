package cmd

import (
	"regexp"

	"github.com/gin-contrib/cors"

	"go-kitboxpro/internal/biz"
	"go-kitboxpro/internal/cron"
	"go-kitboxpro/internal/system"

	"github.com/knadh/koanf/v2"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"go-kitboxpro/internal/routes"
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
	//config cors
	origins := m.conf.Strings("cors.allowOrigin")
	credentials := m.conf.Bool("cors.allowCredentials")
	headers := []string{"Origin", "Content-Type", "Accept", "Authorization"}
	exposeHeaders := []string{"Content-Type", "Authorization", "X-Login"}
	methods := []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	maxAge := m.conf.Duration("cors.maxAge")
	m.logger.Debugf("cors config: origins: %v, credentials: %v, headers: %v, exposeHeaders: %v, methods: %v, maxAge: %v", origins, credentials, headers, exposeHeaders, methods, maxAge)
	var originMap = make(map[string]*regexp.Regexp)
	for i := 0; i < len(origins); i++ {
		re, err := regexp.Compile(origins[i])
		if err != nil {
			m.logger.Warnf("cors.allowOrigin[%s] is invalid, skip it", origins[i])
			continue
		}
		m.logger.Debugf("cors.allowOrigin[%s] is valid", origins[i])
		originMap[origins[i]] = re
	}

	var c = cors.Config{
		AllowOriginFunc: func(origin string) bool {
			for k, v := range originMap {
				if v.MatchString(origin) {
					return true
				}
				m.logger.Debugf("cors.allowOrigin[%s] is not match origin[%s]", k, origin)
			}
			return false
		},
		AllowMethods:     methods,
		AllowHeaders:     headers,
		AllowCredentials: credentials,
		ExposeHeaders:    exposeHeaders,
		MaxAge:           maxAge,
	}

	m.http.ConfigCors(c)
	//config cors end
}

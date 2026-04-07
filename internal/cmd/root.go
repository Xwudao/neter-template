package cmd

import (
	"github.com/knadh/koanf/v2"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/Xwudao/neter-template/internal/biz"
	"github.com/Xwudao/neter-template/internal/cron"
	"github.com/Xwudao/neter-template/internal/system"

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
	m.http.Register()
	return m.http.Run()
}

// check system
func (m *MainApp) checkSystem() {
	if err := m.initSystem.CheckSystem(); err != nil {
		panic(err)
	}
}

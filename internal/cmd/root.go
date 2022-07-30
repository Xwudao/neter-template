/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"regexp"

	"github.com/knadh/koanf"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/Xwudao/neter-template/internal/core"
	"github.com/Xwudao/neter-template/internal/routes"
	"github.com/Xwudao/neter-template/pkg/utils/cron"
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
	initSystem *core.InitSystem
	cron       *cron.Cron
	conf       *koanf.Koanf
	logger     *zap.SugaredLogger
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
			m.logger.Warnf("cors.allow_origin[%s] is invalid, skip it", origins[i])
			continue
		}
		m.logger.Debugf("cors.allow_origin[%s] is valid", origins[i])
		originMap[origins[i]] = re
	}

	m.http.SetOriginFun(func(origin string) bool {
		for k, v := range originMap {
			if v.MatchString(origin) {
				return true
			}
			m.logger.Debugf("cors.allow_origin[%s] is not match origin[%s]", k, origin)
		}
		return false
	})
	m.http.SetCredentials(credentials)
	m.http.SetHeaders(headers)
	m.http.SetExposeHeaders(exposeHeaders)
	m.http.SetMethods(methods)
	m.http.SetMaxAge(maxAge)
	m.http.ConfigCors()
	//config cors end
}

func (m *MainApp) Run() error {
	m.initSystem.InitConfig()
	m.cron.Run()
	m.cors() //this must be set before Register method
	m.http.Register()
	return m.http.Run()
}

func NewMainApp(http *routes.HttpEngine, logger *zap.SugaredLogger, conf *koanf.Koanf, cron *cron.Cron, initSystem *core.InitSystem) (*MainApp, func()) {
	m := &MainApp{
		logger:     logger,
		http:       http,
		initSystem: initSystem,
		cron:       cron,
		conf:       conf,
	}
	cleanup := func() {
		_ = m.cron.Close()
	}

	return m, cleanup
}

//func Execute() {
//	err := rootCmd.Execute()
//	if err != nil {
//		os.Exit(1)
//	}
//}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.internal.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

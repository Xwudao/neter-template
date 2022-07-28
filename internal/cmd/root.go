/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/knadh/koanf"
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

func NewApp(r *gin.Engine, log *zap.SugaredLogger, routes *routes.AppRoutes, conf *koanf.Koanf) *cobra.Command {
	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		routes.Setup()

		port := conf.Int("app.port")
		err := r.Run(fmt.Sprintf(":%d", port))
		if err != nil {
			panic(err)
		}
		log.Infof("app running on port: %d", port)
	}

	return rootCmd
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

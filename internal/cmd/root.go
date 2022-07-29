/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/spf13/cobra"

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
	http *routes.HttpEngine
}

func (m *MainApp) Run() error {
	m.http.Register()
	return m.http.Run()
}

func NewMainApp(http *routes.HttpEngine) *MainApp {
	return &MainApp{http: http}
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

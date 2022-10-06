/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/Xwudao/neter-template/internal/cmd_app"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init something that the system need",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
	},
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "config something that the system need",
	Run: func(cmd *cobra.Command, args []string) {
		ic, f, err := cmd_app.InitCmd()
		if err != nil {
			panic(err)
		}
		defer f()

		ic.Config()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

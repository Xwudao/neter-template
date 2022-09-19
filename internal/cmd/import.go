/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/Xwudao/neter-template/internal/core"

	"github.com/spf13/cobra"
)

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import",
	Short: "this is command is an example",
	Long:  `you can learn how to use dependency in sub command.`,
	Run: func(cmd *cobra.Command, args []string) {
		f, clear, err := core.CmdApp()
		if err != nil {
			panic("init dependency error")
		}

		f.Log.Infof("importing...")
		f.Koanf.Print()

		defer clear()

	},
}

func init() {
	// rootCmd.AddCommand(importCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// importCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// importCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

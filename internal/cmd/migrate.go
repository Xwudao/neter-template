/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"

	"github.com/spf13/cobra"

	"github.com/Xwudao/neter-template/internal/cmd_app"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate database",
	Run: func(cmd *cobra.Command, args []string) {
		migrateName, _ := cmd.Flags().GetString("name")
		mc, lifecycle, err := cmd_app.MigrateCmd()
		if err != nil {
			panic(err)
		}
		defer func() { _ = lifecycle.Stop(context.Background()) }()

		mc.Run(migrateName)

	},
}

// up command

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "up database",
	Run: func(cmd *cobra.Command, args []string) {
		upAll, _ := cmd.Flags().GetBool("all")

		mc, lifecycle, err := cmd_app.MigrateCmd()
		if err != nil {
			panic(err)
		}
		defer func() { _ = lifecycle.Stop(context.Background()) }()

		if err := mc.Up(upAll); err != nil {
			panic(err)
		}
	},
}

var downCmd = &cobra.Command{
	Use:   "down",
	Short: "down database",
	Run: func(cmd *cobra.Command, args []string) {
		downAll, _ := cmd.Flags().GetBool("all")

		mc, lifecycle, err := cmd_app.MigrateCmd()
		if err != nil {
			panic(err)
		}
		defer func() { _ = lifecycle.Stop(context.Background()) }()

		if err := mc.Down(downAll); err != nil {
			panic(err)
		}
	},
}

func init() {
	migrateCmd.AddCommand(upCmd, downCmd)
	rootCmd.AddCommand(migrateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// migrateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// migrateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	migrateCmd.Flags().StringP("name", "n", "", "the migration name for this migrate.")
	// _ = migrateCmd.MarkFlagRequired("name")

	upCmd.Flags().BoolP("all", "a", false, "up all migrations")

	downCmd.Flags().BoolP("all", "a", false, "down all migrations")
}

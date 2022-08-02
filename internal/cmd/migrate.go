/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"

	"ariga.io/atlas/sql/sqltool"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"

	"github.com/Xwudao/neter-template/internal/core"
	"github.com/Xwudao/neter-template/internal/data/ent/migrate"

	gomigrate "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate database",
	Run: func(cmd *cobra.Command, args []string) {
		app, cleanup, err := core.CmdApp()
		if err != nil {
			fmt.Println(err)
			return
		}
		conf := app.Conf
		log := app.Log

		dsn := conf.String("db.devDsn")

		migrateName, _ := cmd.Flags().GetString("name")
		// dial := conf.String("db.dialect")

		// dir, err := sqltool.NewGolangMigrateDir("migrations")
		// if err != nil {
		// 	log.Fatalf("failed creating atlas migration directory: %v", err)
		// }

		ctx := context.Background()
		// Create a local migration directory able to understand golang-migrate migration files for replay.
		dir, err := sqltool.NewGolangMigrateDir("migrations")
		if err != nil {
			log.Fatalf("failed creating atlas migration directory: %v", err)
		}
		// Write migration diff.
		opts := []schema.MigrateOption{
			schema.WithDir(dir),                         // provide migration directory
			schema.WithMigrationMode(schema.ModeReplay), // provide migration mode
			schema.WithDialect(dialect.MySQL),           // Ent dialect to use
			schema.WithFormatter(sqltool.GolangMigrateFormatter),
		}
		if migrateName == "" {
			err = migrate.Diff(ctx, dsn, opts...)
		} else {
			err = migrate.NamedDiff(ctx, dsn, migrateName, opts...)
		}
		if err != nil {
			log.Fatalf("failed generating migration file: %v", err)
		}

		defer cleanup()

	},
}

// up command

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "up database",
	Run: func(cmd *cobra.Command, args []string) {
		app, cleanup, err := core.CmdApp()
		if err != nil {
			panic(err)
		}
		defer cleanup()

		upAll, _ := cmd.Flags().GetBool("all")

		conf := app.Conf
		log := app.Log
		dialect := conf.String("db.dialect")
		host := conf.String("db.host")
		port := conf.String("db.port")
		username := conf.String("db.username")
		password := conf.String("db.password")
		database := conf.String("db.database")

		dsn := fmt.Sprintf("%s://%s:%s@tcp(%s:%s)/%s?multiStatements=true", dialect, username, password, host, port, database)

		m, err := gomigrate.New(fmt.Sprintf("file://%s", "migrations"), dsn)
		if err != nil {
			log.Fatalf("failed creating migrate: %v", err)
		}

		if upAll {
			err = m.Up()
		} else {
			err = m.Steps(1)
		}
		if err != nil {
			log.Errorf("failed creating migrate: %v", err)
			return
		}

		log.Infof("migration up success")

	},
}

var downCmd = &cobra.Command{
	Use:   "down",
	Short: "down database",
	Run: func(cmd *cobra.Command, args []string) {
		app, cleanup, err := core.CmdApp()
		if err != nil {
			panic(err)
		}
		defer cleanup()

		downAll, _ := cmd.Flags().GetBool("all")

		conf := app.Conf
		log := app.Log
		dialect := conf.String("db.dialect")
		host := conf.String("db.host")
		port := conf.String("db.port")
		username := conf.String("db.username")
		password := conf.String("db.password")
		database := conf.String("db.database")

		dsn := fmt.Sprintf("%s://%s:%s@tcp(%s:%s)/%s?multiStatements=true", dialect, username, password, host, port, database)

		m, err := gomigrate.New(fmt.Sprintf("file://%s", "migrations"), dsn)
		if err != nil {
			log.Fatalf("failed creating migrate: %v", err)
		}

		if downAll {
			err = m.Down()
		} else {
			err = m.Steps(-1)
		}
		if err != nil {
			log.Errorf("failed creating migrate: %v", err)
			return
		}

		log.Infof("migration down success")
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
	//_ = migrateCmd.MarkFlagRequired("name")

	upCmd.Flags().BoolP("all", "a", false, "up all migrations")

	downCmd.Flags().BoolP("all", "a", false, "down all migrations")
}

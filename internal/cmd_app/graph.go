package cmd_app

import (
	"github.com/Xwudao/loom"

	"github.com/Xwudao/neter-template/internal/system"
	"github.com/Xwudao/neter-template/pkg/config"
	"github.com/Xwudao/neter-template/pkg/logger"
)

//go:generate go tool loom generate .

var (
	// migrateAppGraph keeps the name MigrateCmd so internal/cmd/migrate.go is
	// unchanged by the Wire → Loom migration.
	migrateAppGraph = loom.Graph[*MigrateApp](
		loom.Name("MigrateCmd"),
		loom.Provide(NewMigrateApp),
		loom.Provide(system.NewAppContext),
		loom.Provide(logger.NewLogger),
		loom.Provide(config.NewKoanf),
	)

	// initAppGraph keeps the name InitCmd for internal/cmd/init.go.
	initAppGraph = loom.Graph[*InitApp](
		loom.Name("InitCmd"),
		loom.Provide(NewInitApp),
	)
)

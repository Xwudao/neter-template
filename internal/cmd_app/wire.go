//go:build wireinject
// +build wireinject

package cmd_app

import (
	"github.com/google/wire"
	"go-kitboxpro/internal/system"
	"go-kitboxpro/pkg/config"
	"go-kitboxpro/pkg/logger"
)

func MigrateCmd() (*MigrateApp, func(), error) {
	panic(wire.Build(
		NewMigrateApp,
		system.NewAppContext,
		logger.NewLogger,
		config.NewKoanf,
	))
}

func InitCmd() (*InitApp, func(), error) {
	panic(wire.Build(NewInitApp))
}

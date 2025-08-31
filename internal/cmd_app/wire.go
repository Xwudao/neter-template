//go:build wireinject
// +build wireinject

package cmd_app

import (
	"github.com/Xwudao/neter-template/internal/system"
	"github.com/Xwudao/neter-template/pkg/config"
	"github.com/Xwudao/neter-template/pkg/logger"
	"github.com/google/wire"
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

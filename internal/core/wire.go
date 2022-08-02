//go:build wireinject
// +build wireinject

package core

import (
	"github.com/google/wire"

	"github.com/Xwudao/neter-template/pkg/config"
	"github.com/Xwudao/neter-template/pkg/logger"
)

func CmdApp() (*App, func(), error) {
	panic(wire.Build(NewApp, logger.NewLogger, config.NewConfig))
}

func TestApp() (*TestManager, func(), error) {
	panic(wire.Build(
		NewTestApp,
		NewTestAppContext,
		config.NewTestConfig,
		logger.NewLogger,
	))
}

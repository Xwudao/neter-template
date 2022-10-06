//go:build wireinject
// +build wireinject

package core

import (
	"github.com/Xwudao/neter-template/internal/system"
	"github.com/google/wire"

	"github.com/Xwudao/neter-template/pkg/config"
	"github.com/Xwudao/neter-template/pkg/logger"
)

// func CmdApp() (*App, func(), error) {
// 	panic(wire.Build(
// 		NewApp,
// 		logger.NewLogger,
// 		config.NewKoanf,
// 		config.NewConfig,
// 	))
// }

func TestApp() (*Test, func(), error) {
	panic(wire.Build(
		NewTestApp,
		system.NewTestAppContext,
		config.NewTestConfig,
		logger.NewTestLogger,
	))
}

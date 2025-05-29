//go:build wireinject
// +build wireinject

package core

import (
	"github.com/google/wire"
	"go-kitboxpro/internal/system"

	"go-kitboxpro/pkg/config"
	"go-kitboxpro/pkg/logger"
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

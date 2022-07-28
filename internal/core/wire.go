//go:build wireinject
// +build wireinject

package core

import (
	"github.com/Xwudao/neter-template/pkg/config"
	"github.com/Xwudao/neter-template/pkg/logger"
	"github.com/google/wire"
)

func CoreApp() (*App, func(), error) {
	panic(wire.Build(NewApp, logger.NewLogger, config.NewConfig))
}

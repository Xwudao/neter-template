//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"

	"github.com/Xwudao/neter-template/internal/biz"
	"github.com/Xwudao/neter-template/internal/cmd"
	"github.com/Xwudao/neter-template/internal/routes"
	"github.com/Xwudao/neter-template/pkg/config"
	"github.com/Xwudao/neter-template/pkg/logger"
)

func mainApp() (*cmd.MainApp, func(), error) {
	panic(wire.Build(cmd.NewMainApp, config.NewConfig, logger.NewLogger, biz.ProviderBizSet, routes.ProviderRouteSet))
}

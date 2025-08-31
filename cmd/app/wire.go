//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"

	"github.com/Xwudao/neter-template/internal/cron"
	"github.com/Xwudao/neter-template/internal/data"
	"github.com/Xwudao/neter-template/internal/system"
	"github.com/Xwudao/neter-template/pkg/utils"

	"github.com/Xwudao/neter-template/internal/biz"
	"github.com/Xwudao/neter-template/internal/cmd"
	"github.com/Xwudao/neter-template/internal/core"
	"github.com/Xwudao/neter-template/internal/routes"
	"github.com/Xwudao/neter-template/pkg/config"
	"github.com/Xwudao/neter-template/pkg/logger"
)

func mainApp() (*cmd.MainApp, func(), error) {
	panic(wire.Build(
		cmd.NewMainApp,
		logger.NewLogger,
		logger.NewZapWriter,
		config.ProviderConfigSet,
		cron.ProviderCronSet,
		core.ProviderCoreSet,
		biz.ProviderBizSet,
		data.ProviderDataSet,
		utils.ProvideUtilSet,
		routes.ProviderRouteSet,
		system.ProviderSystemSet,
	))
}

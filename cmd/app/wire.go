//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"

	"go-kitboxpro/internal/cron"
	"go-kitboxpro/internal/data"
	"go-kitboxpro/internal/system"
	"go-kitboxpro/pkg/utils"

	"go-kitboxpro/internal/biz"
	"go-kitboxpro/internal/cmd"
	"go-kitboxpro/internal/core"
	"go-kitboxpro/internal/routes"
	"go-kitboxpro/pkg/config"
	"go-kitboxpro/pkg/logger"
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

//go:build wireinject
// +build wireinject

package cmd

import (
	"github.com/Xwudao/neter-template/internal/system"
	"github.com/Xwudao/neter-template/pkg/utils/cron"
	"github.com/google/wire"

	"github.com/Xwudao/neter-template/internal/biz"
	"github.com/Xwudao/neter-template/internal/core"
	"github.com/Xwudao/neter-template/internal/routes"
	"github.com/Xwudao/neter-template/pkg/config"
	"github.com/Xwudao/neter-template/pkg/logger"
)

func MainApp() (*App, func(), error) {
	panic(wire.Build(
		NewApp,
		config.NewKoanf,
		config.NewConfig,
		logger.NewLogger,
		cron.ProviderCronSet,
		core.ProviderCoreSet,
		biz.ProviderBizSet,
		routes.ProviderRouteSet,
		system.ProviderSystemSet,
		ProvideCmdSet,
	))
}

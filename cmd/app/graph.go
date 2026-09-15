package main

import (
	"github.com/Xwudao/loom"

	"github.com/Xwudao/neter-template/internal/biz"
	"github.com/Xwudao/neter-template/internal/cmd"
	"github.com/Xwudao/neter-template/internal/cmd_app"
	"github.com/Xwudao/neter-template/internal/core"
	"github.com/Xwudao/neter-template/internal/cron"
	"github.com/Xwudao/neter-template/internal/data"
	"github.com/Xwudao/neter-template/internal/routes"
	"github.com/Xwudao/neter-template/internal/system"
	"github.com/Xwudao/neter-template/pkg/config"
	"github.com/Xwudao/neter-template/pkg/logger"
	"github.com/Xwudao/neter-template/pkg/utils"
)

//go:generate go tool loom generate .

// mainAppGraph is the whole-application dependency graph.
//
// loom.Name keeps the generated constructor called mainApp, so main.go keeps
// the same call it had with the Wire injector it replaces.
var mainAppGraph = loom.Graph[*cmd.MainApp](
	loom.Name("mainApp"),
	loom.Provide(cmd.NewMainApp),
	loom.Provide(cmd_app.NewMigrateApp),
	loom.Provide(logger.NewLogger),
	loom.Provide(logger.NewZapWriter),
	config.ProviderConfigSet,
	cron.ProviderCronSet,
	core.ProviderCoreSet,
	biz.ProviderBizSet,
	data.ProviderDataSet,
	utils.ProvideUtilSet,
	routes.ProviderRouteSet,
	system.ProviderSystemSet,
)

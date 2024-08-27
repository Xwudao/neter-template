// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/Xwudao/neter-template/internal/biz"
	"github.com/Xwudao/neter-template/internal/cmd"
	"github.com/Xwudao/neter-template/internal/cron"
	"github.com/Xwudao/neter-template/internal/data"
	"github.com/Xwudao/neter-template/internal/routes"
	"github.com/Xwudao/neter-template/internal/routes/v1"
	"github.com/Xwudao/neter-template/internal/system"
	"github.com/Xwudao/neter-template/pkg/config"
	"github.com/Xwudao/neter-template/pkg/logger"
	"github.com/Xwudao/neter-template/pkg/utils/jwt"
)

// Injectors from wire.go:

func mainApp() (*cmd.MainApp, func(), error) {
	koanf, err := config.NewKoanf()
	if err != nil {
		return nil, nil, err
	}
	sugaredLogger, err := logger.NewLogger(koanf)
	if err != nil {
		return nil, nil, err
	}
	zapWriter := logger.NewZapWriter(sugaredLogger)
	jwtConfig, err := config.NewJwtConfigData(koanf)
	if err != nil {
		return nil, nil, err
	}
	client := jwt.NewClient(jwtConfig)
	dbConfig, err := config.NewDBConfig(koanf)
	if err != nil {
		return nil, nil, err
	}
	dataData, err := data.NewData(koanf, dbConfig)
	if err != nil {
		return nil, nil, err
	}
	engine := routes.NewEngine(zapWriter, client, dataData, koanf, sugaredLogger)
	appContext := system.NewAppContext()
	userRepository := data.NewUserRepository(appContext, dataData)
	userBiz := biz.NewUserBiz(sugaredLogger, userRepository, client, appContext)
	userRoute := v1.NewUserRoute(engine, userBiz, koanf)
	siteConfigRepository := data.NewSiteConfigRepository(appContext, dataData)
	siteConfigBiz := biz.NewSiteConfigBiz(sugaredLogger, siteConfigRepository, appContext)
	siteHelpBiz := biz.NewSiteHelpBiz(sugaredLogger, siteConfigBiz, appContext)
	siteConfigRoute := v1.NewSiteConfigRoute(engine, siteConfigBiz, siteHelpBiz, sugaredLogger, koanf)
	dataListRepository := data.NewDataListRepository(appContext, dataData)
	dataListBiz := biz.NewDataListBiz(sugaredLogger, dataListRepository, appContext)
	dataListRoute := v1.NewDataListRoute(engine, dataListBiz, sugaredLogger, koanf)
	htmlHelpBiz := biz.NewHtmlHelpBiz(sugaredLogger, appContext, siteConfigBiz, dataListBiz, client)
	htmlRoute := v1.NewHtmlRoute(engine, htmlHelpBiz, sugaredLogger, koanf)
	httpEngine, err := routes.NewHttpEngine(engine, koanf, sugaredLogger, appContext, userRoute, siteConfigRoute, dataListRoute, htmlRoute)
	if err != nil {
		return nil, nil, err
	}
	cronCron, err := cron.NewCron(sugaredLogger)
	if err != nil {
		return nil, nil, err
	}
	initSystem := system.NewInitSystem(koanf)
	systemInitBiz := biz.NewSystemInitBiz(sugaredLogger, userRepository, appContext)
	cmdMainApp, cleanup := cmd.NewMainApp(httpEngine, sugaredLogger, koanf, cronCron, initSystem, systemInitBiz, siteConfigBiz)
	return cmdMainApp, func() {
		cleanup()
	}, nil
}

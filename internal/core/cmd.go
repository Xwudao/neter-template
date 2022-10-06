package core

import (
	"github.com/Xwudao/neter-template/internal/system"
	"github.com/knadh/koanf"
	"go.uber.org/zap"
)

// type App struct {
// 	Log    *zap.SugaredLogger
// 	Koanf  *koanf.Koanf
// 	Config *config.AppConfig
// }
//
// func NewApp(
// 	log *zap.SugaredLogger,
// 	config *config.AppConfig,
// 	conf *koanf.Koanf,
// ) *App {
// 	return &App{
// 		Log:    log,
// 		Koanf:  conf,
// 		Config: config,
// 	}
// }

// Test for test env
type Test struct {
	Logger *zap.SugaredLogger
	Conf   *koanf.Koanf
	Ctx    *system.AppContext
}

func NewTestApp(logger *zap.SugaredLogger, ctx *system.AppContext, conf *koanf.Koanf) *Test {
	return &Test{
		Logger: logger,
		Conf:   conf,
		Ctx:    ctx,
	}
}

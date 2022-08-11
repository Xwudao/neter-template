package core

import (
	"github.com/knadh/koanf"
	"go.uber.org/zap"
)

type App struct {
	Log  *zap.SugaredLogger
	Conf *koanf.Koanf
}

func NewApp(log *zap.SugaredLogger, conf *koanf.Koanf) *App {
	return &App{Log: log, Conf: conf}
}

// Test for test env
type Test struct {
	Logger *zap.SugaredLogger
	Conf   *koanf.Koanf
	Ctx    *AppContext
}

func NewTestApp(logger *zap.SugaredLogger, ctx *AppContext, conf *koanf.Koanf) *Test {
	return &Test{
		Logger: logger,
		Conf:   conf,
		Ctx:    ctx,
	}
}

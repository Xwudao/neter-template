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

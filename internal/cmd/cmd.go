package cmd

import (
	"github.com/knadh/koanf"
	"go.uber.org/zap"
)

type App struct {
	log  *zap.SugaredLogger
	conf *koanf.Koanf
}

func NewApp(log *zap.SugaredLogger, conf *koanf.Koanf) *App {
	return &App{log: log, conf: conf}
}

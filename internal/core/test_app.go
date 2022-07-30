package core

import (
	"github.com/knadh/koanf"
	"go.uber.org/zap"
)

type TestManager struct {
	Logger *zap.SugaredLogger
	Conf   *koanf.Koanf
	Ctx    *AppContext
}

func NewTestApp(logger *zap.SugaredLogger, ctx *AppContext, conf *koanf.Koanf) *TestManager {
	return &TestManager{
		Logger: logger,
		Conf:   conf,
		Ctx:    ctx,
	}
}

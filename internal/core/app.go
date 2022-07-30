package core

import (
	"context"
)

type AppContext struct {
	Ctx    context.Context
	Cancel context.CancelFunc
}

func NewAppContext() *AppContext {
	ctx, cancel := context.WithCancel(context.Background())
	return &AppContext{
		Ctx:    ctx,
		Cancel: cancel,
	}
}

func NewTestAppContext() *AppContext {
	return NewAppContext()
}

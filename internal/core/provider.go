package core

import (
	"github.com/google/wire"
	"go-kitboxpro/internal/system"
)

var ProviderCoreSet = wire.NewSet(system.NewAppContext)

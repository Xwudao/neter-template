package core

import (
	"github.com/Xwudao/neter-template/internal/system"
	"github.com/google/wire"
)

var ProviderCoreSet = wire.NewSet(system.NewAppContext)

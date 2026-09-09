package core

import (
	"github.com/google/wire"
	"github.com/Xwudao/neter-template/internal/system"
)

var ProviderCoreSet = wire.NewSet(system.NewAppContext)

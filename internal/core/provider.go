package core

import (
	"github.com/Xwudao/loom"

	"github.com/Xwudao/neter-template/internal/system"
)

var ProviderCoreSet = loom.Module(loom.Provide(system.NewAppContext))

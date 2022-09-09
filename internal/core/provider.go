package core

import (
	"github.com/google/wire"
)

var ProviderCoreSet = wire.NewSet(NewAppContext)

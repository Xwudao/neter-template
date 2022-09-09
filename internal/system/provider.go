package system

import (
	"github.com/google/wire"
)

var ProviderSystemSet = wire.NewSet(NewInitSystem)

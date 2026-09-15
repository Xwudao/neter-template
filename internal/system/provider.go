package system

import "github.com/Xwudao/loom"

var ProviderSystemSet = loom.Module(loom.Provide(NewInitSystem))

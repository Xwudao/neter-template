package data

import "github.com/Xwudao/loom"

var ProviderDataSet = loom.Module(
	loom.Provide(NewData),
	loom.Provide(NewUserRepository),
	loom.Provide(NewSiteConfigRepository),
	loom.Provide(NewDataListRepository),
)

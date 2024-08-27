package data

import "github.com/google/wire"

var ProviderDataSet = wire.NewSet(
	NewData,
	NewUserRepository,
	NewSiteConfigRepository,
	NewDataListRepository,
)

//var ProviderDataSet = wire.NewSet(NewRedisClient, NewUserRepository)

package data

import "github.com/google/wire"

var ProviderDataSet = wire.NewSet(NewRedisClient)

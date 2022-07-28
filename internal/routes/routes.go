package routes

import (
	"github.com/google/wire"
)

var ProviderRouteSet = wire.NewSet(NewHttpEngine)

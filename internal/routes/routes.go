package routes

import (
	"github.com/google/wire"

	v1 "github.com/Xwudao/neter/internal/routes/v1"
)

var ProviderRouteSet = wire.NewSet(NewEngine, v1.NewHomeRoute, NewAppRoutes)

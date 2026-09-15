package routes

import (
	"github.com/Xwudao/loom"

	v1 "github.com/Xwudao/neter-template/internal/routes/v1"
)

var ProviderRouteSet = loom.Module(
	loom.Provide(NewEngine),
	loom.Provide(NewHttpEngine),
	loom.Provide(NewRouteRegistry),
	loom.Provide(v1.NewUserRoute),
	loom.Provide(v1.NewSiteConfigRoute),
	loom.Provide(v1.NewDataListRoute),
)

package routes

import (
	"github.com/google/wire"

	v1 "github.com/Xwudao/neter-template/internal/routes/v1"
)

var ProviderRouteSet = wire.NewSet(
	NewEngine,
	NewHttpEngine,
	v1.NewUserRoute,
	v1.NewSiteConfigRoute,
	v1.NewDataListRoute,
	v1.NewHtmlRoute,
)

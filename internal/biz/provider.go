package biz

import "github.com/google/wire"

var ProviderBizSet = wire.NewSet(
	NewUserBiz,
	NewSystemInitBiz,
	NewSiteConfigBiz,
	NewSiteHelpBiz,
	NewDataListBiz,
	NewHtmlHelpBiz,
	NewSeoBizBiz,
	// Bind concrete implementations to their handler-facing interfaces.
	// This lets Wire inject the interface type into route constructors.
	wire.Bind(new(UserBizIface), new(*UserBiz)),
	wire.Bind(new(SiteConfigBizIface), new(*SiteConfigBiz)),
	wire.Bind(new(SiteHelpBizIface), new(*SiteHelpBiz)),
	wire.Bind(new(DataListBizIface), new(*DataListBiz)),
	wire.Bind(new(HtmlHelpBizIface), new(*HtmlHelpBiz)),
)

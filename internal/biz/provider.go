package biz

import "github.com/Xwudao/loom"

var ProviderBizSet = loom.Module(
	loom.Provide(NewSystemInitBiz),
	loom.Provide(NewSeoBizBiz),
	// Bind concrete implementations to their handler-facing interfaces.
	// loom.As replaces the Wire Provide+Bind pair: it registers the
	// constructor and exposes the same instance as the interface, so the
	// constructor must not also appear as a plain loom.Provide.
	loom.As[UserBizIface](NewUserBiz),
	loom.As[SiteConfigBizIface](NewSiteConfigBiz),
	loom.As[SiteHelpBizIface](NewSiteHelpBiz),
	loom.As[DataListBizIface](NewDataListBiz),
	loom.As[HtmlHelpBizIface](NewHtmlHelpBiz),
)

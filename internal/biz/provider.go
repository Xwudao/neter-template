package biz

import "github.com/google/wire"

var ProviderBizSet = wire.NewSet(
	NewUserBiz,
	NewSystemInitBiz,
	NewSiteConfigBiz,
	NewSiteHelpBiz,
	NewDataListBiz,
	NewHtmlHelpBiz,
)

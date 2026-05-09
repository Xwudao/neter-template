package biz

import "context"

// SiteHelpBizIface is the interface consumed by route handlers.
type SiteHelpBizIface interface {
	GenerateSiteMap(ctx context.Context, mapPath string) error
}

// Compile-time assertion: *SiteHelpBiz must satisfy SiteHelpBizIface.
var _ SiteHelpBizIface = (*SiteHelpBiz)(nil)

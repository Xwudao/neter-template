package biz

import (
	"context"
	"os"

	"github.com/snabb/sitemap"
	"go.uber.org/zap"

	"go-kitboxpro/internal/domain/models"
	"go-kitboxpro/internal/system"
	"go-kitboxpro/pkg/enum"
)

// SiteHelpBiz sitemap, site helper
type SiteHelpBiz struct {
	log    *zap.SugaredLogger
	appCtx *system.AppContext
	scb    *SiteConfigBiz
}

func NewSiteHelpBiz(log *zap.SugaredLogger, scb *SiteConfigBiz, appCtx *system.AppContext) *SiteHelpBiz {
	return &SiteHelpBiz{
		log:    log.Named("site-help-biz"),
		appCtx: appCtx,
		scb:    scb,
	}
}

func (h *SiteHelpBiz) GenerateSiteMap(ctx context.Context, mapPath string) error {
	sm := sitemap.New()
	//var page = 1
	var siteInfo models.SiteInfoConfig
	if err := h.scb.GetConfig(enum.ConfigKeySiteInfo, &siteInfo); err != nil {
		return err
	}

	// logic for gen sitemap

	f, err := os.OpenFile(mapPath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = sm.WriteTo(f)
	return err
}

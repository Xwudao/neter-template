package biz

import (
	"context"

	"github.com/Xwudao/neter-template/internal/data/sqlc"
	"github.com/Xwudao/neter-template/internal/domain/params"
	"github.com/Xwudao/neter-template/pkg/enum"
)

// SiteConfigBizIface is the interface consumed by route handlers.
type SiteConfigBizIface interface {
	Init() error
	GetConfig(key enum.ConfigKey, target any) error
	UpdateConfig(ctx context.Context, config *params.UpdateSiteConfigParams) error
	Delete(ctx context.Context, id int64) error
	Get(ctx context.Context, id int64) (*sqlc.SiteConfig, error)
	Create(ctx context.Context, p *params.CreateSiteConfigParams) (*sqlc.SiteConfig, error)
	GetAll(ctx context.Context, isAdmin bool) (map[string]string, error)
}

// Compile-time assertion: *SiteConfigBiz must satisfy SiteConfigBizIface.
var _ SiteConfigBizIface = (*SiteConfigBiz)(nil)

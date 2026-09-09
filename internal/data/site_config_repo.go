package data

import (
	"context"

	"github.com/Xwudao/neter-template/internal/biz"
	"github.com/Xwudao/neter-template/internal/data/sqlc"
	"github.com/Xwudao/neter-template/internal/domain/params"
	"github.com/Xwudao/neter-template/internal/system"
)

var _ biz.SiteConfigRepository = (*siteConfigRepository)(nil)

type siteConfigRepository struct {
	appCtx *system.AppContext
	data   *Data
}

func NewSiteConfigRepository(appCtx *system.AppContext, data *Data) biz.SiteConfigRepository {
	return &siteConfigRepository{appCtx: appCtx, data: data}
}

func (u *siteConfigRepository) GetNames(ctx context.Context) ([]string, error) {
	return u.data.Queries.ListSiteConfigNames(ctx)
}

func (u *siteConfigRepository) GetByName(ctx context.Context, name string) (*sqlc.SiteConfig, error) {
	return u.data.Queries.GetSiteConfigByName(ctx, name)
}

func (u *siteConfigRepository) GetAll(ctx context.Context) ([]*sqlc.SiteConfig, error) {
	return u.data.Queries.ListSiteConfigs(ctx)
}

func (u *siteConfigRepository) DeleteByID(ctx context.Context, id int64) error {
	_, err := u.data.Queries.DeleteSiteConfig(ctx, id)
	return err
}

func (u *siteConfigRepository) GetByID(ctx context.Context, id int64) (*sqlc.SiteConfig, error) {
	return u.data.Queries.GetSiteConfig(ctx, id)
}

func (u *siteConfigRepository) Create(ctx context.Context, p *params.CreateSiteConfigParams) (*sqlc.SiteConfig, error) {
	return u.data.Queries.CreateSiteConfig(ctx, sqlc.CreateSiteConfigParams{Name: p.Name, Config: p.Config})
}

func (u *siteConfigRepository) Update(ctx context.Context, p *params.UpdateSiteConfigParams) (int, error) {
	rows, err := u.data.Queries.UpdateSiteConfig(ctx, sqlc.UpdateSiteConfigParams{Name: p.Name, Config: p.Config})
	return int(rows), err
}

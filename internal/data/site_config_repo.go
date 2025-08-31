package data

import (
	"context"

	"github.com/Xwudao/neter-template/internal/biz"
	"github.com/Xwudao/neter-template/internal/data/ent"
	"github.com/Xwudao/neter-template/internal/data/ent/siteconfig"
	"github.com/Xwudao/neter-template/internal/domain/params"
	"github.com/Xwudao/neter-template/internal/system"
)

var _ biz.SiteConfigRepository = (*siteConfigRepository)(nil)

type siteConfigRepository struct {
	appCtx *system.AppContext
	data   *Data
}

func NewSiteConfigRepository(appCtx *system.AppContext, data *Data) biz.SiteConfigRepository {
	return &siteConfigRepository{
		appCtx: appCtx,
		data:   data,
	}
}

// GetNames 获取所有配置名称
func (u *siteConfigRepository) GetNames(ctx context.Context) ([]string, error) {
	return u.data.Client.SiteConfig.Query().Select(siteconfig.FieldName).Strings(ctx)
}

// GetByName 根据名称获取配置
func (u *siteConfigRepository) GetByName(ctx context.Context, name string) (*ent.SiteConfig, error) {
	return u.data.Client.SiteConfig.Query().Where(siteconfig.Name(name)).Only(ctx)
}

func (u *siteConfigRepository) GetAll(ctx context.Context) ([]*ent.SiteConfig, error) {
	return u.data.Client.SiteConfig.Query().All(ctx)
}

func (u *siteConfigRepository) DeleteByID(ctx context.Context, id int64) error {
	return u.data.Client.SiteConfig.DeleteOneID(id).Exec(ctx)
}

func (u *siteConfigRepository) GetByID(ctx context.Context, id int64) (*ent.SiteConfig, error) {
	return u.data.Client.SiteConfig.Get(ctx, id)
}

func (u *siteConfigRepository) Create(ctx context.Context, p *params.CreateSiteConfigParams) (*ent.SiteConfig, error) {
	return u.data.Client.SiteConfig.Create().SetConfig(p.Config).SetName(p.Name).Save(ctx)
}

func (u *siteConfigRepository) Update(ctx context.Context, p *params.UpdateSiteConfigParams) (int, error) {
	return u.data.Client.SiteConfig.Update().Where(siteconfig.Name(p.Name)).SetConfig(p.Config).Save(ctx)
}

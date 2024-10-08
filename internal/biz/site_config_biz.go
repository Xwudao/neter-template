package biz

import (
	"context"
	"sync"

	"github.com/bytedance/sonic"
	"go.uber.org/zap"

	"github.com/Xwudao/neter-template/internal/data/ent"
	"github.com/Xwudao/neter-template/internal/domain/models"
	"github.com/Xwudao/neter-template/internal/domain/params"
	"github.com/Xwudao/neter-template/internal/system"
	"github.com/Xwudao/neter-template/pkg/enum"
	"github.com/Xwudao/neter-template/pkg/varx"
)

type SiteConfigRepository interface {
	GetAll(ctx context.Context) ([]*ent.SiteConfig, error)
	GetNames(ctx context.Context) ([]string, error)
	GetByName(ctx context.Context, name string) (*ent.SiteConfig, error)
	DeleteByID(ctx context.Context, id int64) error
	GetByID(ctx context.Context, id int64) (*ent.SiteConfig, error)
	Create(ctx context.Context, p *params.CreateSiteConfigParams) (*ent.SiteConfig, error)
	Update(ctx context.Context, p *params.UpdateSiteConfigParams) (int, error)
}

type SiteConfigBiz struct {
	log    *zap.SugaredLogger
	appCtx *system.AppContext
	scr    SiteConfigRepository

	sync.Map
}

func NewSiteConfigBiz(log *zap.SugaredLogger, scr SiteConfigRepository, appCtx *system.AppContext) *SiteConfigBiz {
	return &SiteConfigBiz{
		log:    log.Named("site-config-biz"),
		appCtx: appCtx,
		scr:    scr,
	}
}

// Init 初始化
func (h *SiteConfigBiz) Init() error {
	var initMap = map[enum.ConfigKey]string{
		enum.ConfigKeySiteInfo:  varx.MustMarshalDefault(&models.SiteInfoConfig{}),
		enum.ConfigKeySeoConfig: varx.MustMarshalDefault(&models.SEOConfig{}),
	}
	existedNames, err := h.scr.GetNames(h.appCtx.Ctx)
	if err != nil {
		return err
	}
	var initNames []string
	for k := range initMap {
		initNames = append(initNames, string(k))
	}

	var needCreate = varx.ArrDiff(initNames, existedNames)
	h.log.Infoln("needCreate", needCreate)

	for _, name := range needCreate {
		data, err := h.scr.Create(h.appCtx.Ctx, &params.CreateSiteConfigParams{
			Name:   name,
			Config: initMap[enum.ConfigKey(name)],
		})
		if err != nil {
			return err
		}
		h.Map.Store(enum.ConfigKey(name), data.Config)
	}

	return nil
}

// GetConfig 获取配置
func (h *SiteConfigBiz) GetConfig(key enum.ConfigKey, target any) error {
	if data, ok := h.Map.Load(key); ok {
		//return json.Unmarshal([]byte(data.(string)), target)
		return sonic.UnmarshalString(data.(string), target)
	}

	config, err := h.scr.GetByName(h.appCtx.Ctx, string(key))
	if err != nil {
		return err
	}

	if err := sonic.UnmarshalString(config.Config, target); err != nil {
		return err
	}

	h.Map.Store(key, config.Config)
	return nil
}

// UpdateConfig 更新配置
func (h *SiteConfigBiz) UpdateConfig(ctx context.Context, config *params.UpdateSiteConfigParams) error {
	_, err := h.scr.Update(ctx, config)
	if err != nil {
		return err
	}

	h.Map.Store(enum.ConfigKey(config.Name), config.Config)
	return nil
}

func (h *SiteConfigBiz) Delete(ctx context.Context, id int64) error {
	return h.scr.DeleteByID(ctx, id)
}

func (h *SiteConfigBiz) Get(ctx context.Context, id int64) (*ent.SiteConfig, error) {
	return h.scr.GetByID(ctx, id)
}

func (h *SiteConfigBiz) Create(ctx context.Context, p *params.CreateSiteConfigParams) (*ent.SiteConfig, error) {
	return h.scr.Create(ctx, p)
}

func (h *SiteConfigBiz) GetAll(ctx context.Context) (map[string]string, error) {
	data, err := h.scr.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	var m = make(map[string]string)
	for _, v := range data {
		m[v.Name] = v.Config
	}

	return m, nil
}

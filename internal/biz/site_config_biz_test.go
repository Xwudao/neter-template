package biz_test

import (
	"context"
	"errors"
	"testing"
	"time"

	json "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"

	"github.com/Xwudao/neter-template/internal/biz"
	"github.com/Xwudao/neter-template/internal/biz/mocks"
	"github.com/Xwudao/neter-template/internal/data/ent"
	"github.com/Xwudao/neter-template/internal/domain/models"
	"github.com/Xwudao/neter-template/internal/domain/params"
	"github.com/Xwudao/neter-template/internal/system"
	"github.com/Xwudao/neter-template/pkg/enum"
)

func newTestSiteConfigBiz(t *testing.T, scr biz.SiteConfigRepository) *biz.SiteConfigBiz {
	t.Helper()
	log := zap.NewNop().Sugar()
	appCtx := system.NewTestAppContext()
	return biz.NewSiteConfigBiz(log, scr, appCtx)
}

func mustJSON(v any) string {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return string(b)
}

// TestSiteConfigBiz_Init_CreatesMissingKeys verifies that Init creates all missing config keys.
func TestSiteConfigBiz_Init_CreatesMissingKeys(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepo := mocks.NewMockSiteConfigRepository(ctrl)

	// No config keys exist yet.
	mockRepo.EXPECT().
		GetNames(gomock.Any()).
		Return([]string{}, nil)

	// Both site_info and seo_config should be created.
	mockRepo.EXPECT().
		Create(gomock.Any(), gomock.Any()).
		Return(&ent.SiteConfig{ID: 1}, nil).
		Times(2)

	b := newTestSiteConfigBiz(t, mockRepo)
	err := b.Init()
	assert.NoError(t, err)
}

// TestSiteConfigBiz_Init_SkipsExistingKeys verifies that Init skips keys that already exist.
func TestSiteConfigBiz_Init_SkipsExistingKeys(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepo := mocks.NewMockSiteConfigRepository(ctrl)

	// site_info already exists; only seo_config should be created.
	mockRepo.EXPECT().
		GetNames(gomock.Any()).
		Return([]string{string(enum.ConfigKeySiteInfo)}, nil)

	mockRepo.EXPECT().
		Create(gomock.Any(), gomock.Any()).
		Return(&ent.SiteConfig{ID: 2}, nil).
		Times(1)

	b := newTestSiteConfigBiz(t, mockRepo)
	err := b.Init()
	assert.NoError(t, err)
}

// TestSiteConfigBiz_Init_RepoError verifies that an error from GetNames propagates.
func TestSiteConfigBiz_Init_RepoError(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepo := mocks.NewMockSiteConfigRepository(ctrl)
	mockRepo.EXPECT().
		GetNames(gomock.Any()).
		Return(nil, errors.New("db down"))

	b := newTestSiteConfigBiz(t, mockRepo)
	err := b.Init()
	assert.Error(t, err)
}

// TestSiteConfigBiz_GetConfig_FromRepo verifies GetConfig fetches from repo when cache is cold.
func TestSiteConfigBiz_GetConfig_FromRepo(t *testing.T) {
	ctrl := gomock.NewController(t)

	siteInfo := &models.SiteInfoConfig{SiteName: "测试站点"}
	configJSON := mustJSON(siteInfo)

	mockRepo := mocks.NewMockSiteConfigRepository(ctrl)
	mockRepo.EXPECT().
		GetByName(gomock.Any(), string(enum.ConfigKeySiteInfo)).
		Return(&ent.SiteConfig{
			ID:         1,
			Name:       string(enum.ConfigKeySiteInfo),
			Config:     configJSON,
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
		}, nil)

	b := newTestSiteConfigBiz(t, mockRepo)

	var got models.SiteInfoConfig
	err := b.GetConfig(enum.ConfigKeySiteInfo, &got)

	require.NoError(t, err)
	assert.Equal(t, "测试站点", got.SiteName)
}

// TestSiteConfigBiz_GetConfig_CacheHit verifies GetConfig uses the in-memory cache on second call.
func TestSiteConfigBiz_GetConfig_CacheHit(t *testing.T) {
	ctrl := gomock.NewController(t)

	siteInfo := &models.SiteInfoConfig{SiteName: "缓存站点"}
	configJSON := mustJSON(siteInfo)

	mockRepo := mocks.NewMockSiteConfigRepository(ctrl)
	// GetByName must be called exactly once; the second call should use the cache.
	mockRepo.EXPECT().
		GetByName(gomock.Any(), string(enum.ConfigKeySiteInfo)).
		Return(&ent.SiteConfig{
			ID:     1,
			Name:   string(enum.ConfigKeySiteInfo),
			Config: configJSON,
		}, nil).
		Times(1)

	b := newTestSiteConfigBiz(t, mockRepo)

	// First call – populates the cache.
	var first models.SiteInfoConfig
	require.NoError(t, b.GetConfig(enum.ConfigKeySiteInfo, &first))

	// Second call – must hit the cache, not the repo.
	var second models.SiteInfoConfig
	require.NoError(t, b.GetConfig(enum.ConfigKeySiteInfo, &second))

	assert.Equal(t, first.SiteName, second.SiteName)
}

// TestSiteConfigBiz_GetConfig_RepoError verifies that a repo error propagates to the caller.
func TestSiteConfigBiz_GetConfig_RepoError(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepo := mocks.NewMockSiteConfigRepository(ctrl)
	mockRepo.EXPECT().
		GetByName(gomock.Any(), string(enum.ConfigKeySeoConfig)).
		Return(nil, errors.New("not found"))

	b := newTestSiteConfigBiz(t, mockRepo)

	var cfg models.SEOConfig
	err := b.GetConfig(enum.ConfigKeySeoConfig, &cfg)
	assert.Error(t, err)
}

// TestSiteConfigBiz_UpdateConfig verifies that UpdateConfig delegates to the repo and updates cache.
func TestSiteConfigBiz_UpdateConfig(t *testing.T) {
	ctrl := gomock.NewController(t)

	p := &params.UpdateSiteConfigParams{
		Name:   string(enum.ConfigKeySiteInfo),
		Config: `{"site_name":"updated"}`,
	}

	mockRepo := mocks.NewMockSiteConfigRepository(ctrl)
	mockRepo.EXPECT().
		Update(gomock.Any(), p).
		Return(1, nil)

	b := newTestSiteConfigBiz(t, mockRepo)
	err := b.UpdateConfig(context.Background(), p)
	assert.NoError(t, err)

	// Verify cache was updated: next GetConfig should NOT call the repo again.
	mockRepo.EXPECT().
		GetByName(gomock.Any(), gomock.Any()).
		Times(0)

	var cfg models.SiteInfoConfig
	err = b.GetConfig(enum.ConfigKeySiteInfo, &cfg)
	require.NoError(t, err)
	assert.Equal(t, "updated", cfg.SiteName)
}

// TestSiteConfigBiz_Delete verifies Delete delegates to the repository.
func TestSiteConfigBiz_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepo := mocks.NewMockSiteConfigRepository(ctrl)
	mockRepo.EXPECT().
		DeleteByID(gomock.Any(), int64(10)).
		Return(nil)

	b := newTestSiteConfigBiz(t, mockRepo)
	err := b.Delete(context.Background(), 10)
	assert.NoError(t, err)
}

// TestSiteConfigBiz_GetAll_Admin verifies GetAll returns all configs for admin users.
func TestSiteConfigBiz_GetAll_Admin(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockRepo := mocks.NewMockSiteConfigRepository(ctrl)
	mockRepo.EXPECT().
		GetAll(gomock.Any()).
		Return([]*ent.SiteConfig{
			{ID: 1, Name: string(enum.ConfigKeySiteInfo), Config: `{}`},
			{ID: 2, Name: string(enum.ConfigKeySeoConfig), Config: `{}`},
		}, nil)

	b := newTestSiteConfigBiz(t, mockRepo)
	result, err := b.GetAll(context.Background(), true)
	require.NoError(t, err)
	assert.Len(t, result, 2)
}

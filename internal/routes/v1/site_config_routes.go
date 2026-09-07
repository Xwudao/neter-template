package v1

import (
	"errors"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/knadh/koanf/v2"

	"go.uber.org/zap"

	"github.com/Xwudao/neter-template/internal/biz"
	"github.com/Xwudao/neter-template/internal/core"
	"github.com/Xwudao/neter-template/internal/data/ent/user"
	"github.com/Xwudao/neter-template/internal/domain/params"
	"github.com/Xwudao/neter-template/internal/libx"
	"github.com/Xwudao/neter-template/internal/routes/mdw"
	"github.com/Xwudao/neter-template/pkg/varx"
)

type SiteConfigRoute struct {
	conf *koanf.Koanf
	log  *zap.SugaredLogger

	scb biz.SiteConfigBizIface
	shb biz.SiteHelpBizIface

	sf *libx.StaticFile
}

func NewSiteConfigRoute(
	scb biz.SiteConfigBizIface,
	shb biz.SiteHelpBizIface,
	log *zap.SugaredLogger,
	conf *koanf.Koanf,
) *SiteConfigRoute {
	r := &SiteConfigRoute{
		conf: conf,
		scb:  scb, sf: libx.NewStaticFile(),
		shb: shb,
		log: log.Named("site-config-route"),
	}

	return r
}

func (r *SiteConfigRoute) Register(router gin.IRouter) {
	group := router.Group("/v1/site_config")
	{
		group.GET("/all", core.NoInputE(func(c *gin.Context) (map[string]string, error) { return r.getAll(c, false) }))
	}
	authGroup := router.Group("/auth/v1/site_config").Use(mdw.MustLoginMiddleware())
	{
		_ = authGroup
	}
	adminGroup := router.Group("/admin/v1/site_config").Use(mdw.MustWithRoleMiddleware(user.RoleAdmin))
	{
		adminGroup.GET("/all", core.NoInputE(func(c *gin.Context) (map[string]string, error) { return r.getAll(c, true) }))
		adminGroup.GET("/gen_sitemap", core.NoInputE(r.genSitemap))
		adminGroup.POST("/update", core.JSONE(r.update))
		adminGroup.POST("/write_file", core.JSONE(r.writeFile))
	}
}

func (r *SiteConfigRoute) getAll(c *gin.Context, isAdmin bool) (map[string]string, error) {
	configs, err := r.scb.GetAll(c.Request.Context(), isAdmin)
	if err != nil {
		return nil, err
	}
	return configs, nil
}

func (r *SiteConfigRoute) update(c *gin.Context, pm *params.UpdateSiteConfigParams) (*core.EmptyResponse, error) {
	if err := r.scb.UpdateConfig(c.Request.Context(), pm); err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *SiteConfigRoute) writeFile(c *gin.Context, pm *params.WriteFileParams) (*core.EmptyResponse, error) {
	var allowedFiles = []string{"robots.txt"}
	if !varx.ContainEqual(allowedFiles, pm.Filename) {
		return nil, errors.New("文件名不合法")
	}
	if err := r.sf.WriteFile(pm.Filename, []byte(pm.Data)); err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *SiteConfigRoute) genSitemap(c *gin.Context) (string, error) {
	wd, _ := os.Getwd()
	mapPath := wd + "/web/public/sitemap.xml"
	go func() {
		if err := r.shb.GenerateSiteMap(c, mapPath); err != nil {
			r.log.Errorf("GenSitemap error: %v", err)
		}
	}()
	return "后台生成中", nil
}

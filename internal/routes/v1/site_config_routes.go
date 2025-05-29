package v1

import (
	"errors"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/knadh/koanf/v2"

	"go.uber.org/zap"

	"go-kitboxpro/internal/biz"
	"go-kitboxpro/internal/core"
	"go-kitboxpro/internal/data/ent/user"
	"go-kitboxpro/internal/domain/params"
	"go-kitboxpro/internal/libx"
	"go-kitboxpro/internal/routes/mdw"
	"go-kitboxpro/internal/routes/valid"
	"go-kitboxpro/pkg/varx"
)

type SiteConfigRoute struct {
	conf *koanf.Koanf
	g    *gin.Engine
	log  *zap.SugaredLogger

	scb *biz.SiteConfigBiz
	shb *biz.SiteHelpBiz

	sf *libx.StaticFile
}

func NewSiteConfigRoute(
	g *gin.Engine,
	scb *biz.SiteConfigBiz,
	shb *biz.SiteHelpBiz,
	log *zap.SugaredLogger,
	conf *koanf.Koanf,
) *SiteConfigRoute {
	r := &SiteConfigRoute{
		conf: conf,
		g:    g, scb: scb, sf: libx.NewStaticFile(),
		shb: shb,
		log: log.Named("site-config-route"),
	}

	return r
}

func (r *SiteConfigRoute) Reg() {
	// r.g.GET("/v1/site_config", core.WrapData(r.siteConfig()))

	group := r.g.Group("/v1/site_config")
	{
		group.GET("/all", core.WrapData(r.getAll(false)))
	}
	authGroup := r.g.Group("/auth/v1/site_config").Use(mdw.MustLoginMiddleware())
	{
		// authGroup.GET("/auth", core.WrapData(r.siteConfig()))
		_ = authGroup
	}
	adminGroup := r.g.Group("/admin/v1/site_config").Use(mdw.MustWithRoleMiddleware(user.RoleAdmin))
	{
		adminGroup.GET("/all", core.WrapData(r.getAll(true)))
		adminGroup.GET("/gen_sitemap", core.WrapData(r.genSitemap()))
		adminGroup.POST("/update", core.WrapData(r.update()))
		adminGroup.POST("/write_file", core.WrapData(r.writeFile()))
	}
}

func (r *SiteConfigRoute) siteConfig() core.WrappedHandlerFunc {
	return func(c *gin.Context) (any, *core.RtnStatus) {
		//var (
		//    ctx = c.Request.Context()
		//)
		return "hello", nil
	}
}

func (r *SiteConfigRoute) getAll(isAdmin bool) core.WrappedHandlerFunc {
	return func(c *gin.Context) (any, *core.RtnStatus) {
		var (
			ctx = c.Request.Context()
		)
		configs, err := r.scb.GetAll(ctx, isAdmin)
		if err != nil {
			return nil, core.NewRtnWithErr(err)
		}
		return configs, core.Success
	}
}

func (r *SiteConfigRoute) update() core.WrappedHandlerFunc {
	return func(c *gin.Context) (any, *core.RtnStatus) {
		var (
			ctx = c.Request.Context()
			pm  params.UpdateSiteConfigParams
		)

		if err := c.ShouldBindJSON(&pm); err != nil {
			return nil, core.NewRtnWithErr(err)
		}

		if err := r.scb.UpdateConfig(ctx, &pm); err != nil {
			return nil, core.NewRtnWithErr(err)
		}

		return nil, core.Success
	}
}

func (r *SiteConfigRoute) writeFile() core.WrappedHandlerFunc {
	return func(c *gin.Context) (any, *core.RtnStatus) {
		var (
			//ctx = c.Request.Context()
			pm params.WriteFileParams
		)

		if err := c.ShouldBindJSON(&pm); err != nil {
			return nil, core.NewRtnWithErr(valid.GetErrorMsg(&pm, err))
		}

		var allowedFiles = []string{
			"robots.txt",
		}

		if !varx.ContainEqual(allowedFiles, pm.Filename) {
			return nil, core.NewRtnWithErr(errors.New("文件名不合法"))
		}

		if err := r.sf.WriteFile(pm.Filename, []byte(pm.Data)); err != nil {
			return nil, core.NewRtnWithErr(err)
		}

		return nil, core.Success
	}
}

func (r *SiteConfigRoute) genSitemap() core.WrappedHandlerFunc {
	return func(c *gin.Context) (any, *core.RtnStatus) {
		//var ctx = c.Request.Context()

		wd, _ := os.Getwd()
		mapPath := wd + "/web/public/sitemap.xml"

		go func() {
			if err := r.shb.GenerateSiteMap(c, mapPath); err != nil {
				r.log.Errorf("GenSitemap error: %v", err)
				return
			}
		}()

		return "后台生成中", core.Success
	}
}

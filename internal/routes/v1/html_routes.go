package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/knadh/koanf/v2"

	"go.uber.org/zap"

	"github.com/Xwudao/neter-template/internal/biz"
)

type HtmlRoute struct {
	conf *koanf.Koanf
	g    *gin.Engine
	log  *zap.SugaredLogger

	hhb *biz.HtmlHelpBiz
}

func NewHtmlRoute(g *gin.Engine, hhb *biz.HtmlHelpBiz, log *zap.SugaredLogger, conf *koanf.Koanf) *HtmlRoute {
	r := &HtmlRoute{
		conf: conf,
		g:    g, hhb: hhb,
		log: log.Named("html-route"),
	}

	return r
}

func (r *HtmlRoute) Reg() {
	// r.g.GET("/v1/html", core.WrapData(r.html()))

	group := r.g.Group("/")
	{
		group.GET("/", r.index())
	}

}

func (r *HtmlRoute) index() gin.HandlerFunc {
	return func(c *gin.Context) {
		var m, err = r.hhb.BuildIndexMap(c)
		if err != nil {
			r.log.Errorw("BuildIndexMap error", "error", err)
			c.AbortWithStatus(500)
			return
		}
		c.HTML(200, "pages/index.tpl", m)
	}
}

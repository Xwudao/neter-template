package v1

import (
	"net/http"

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
	r.g.LoadHTMLGlob("templates/*.html")
	r.g.Static("/static", "./templates/static")

	r.g.GET("/", func(c *gin.Context) {
		md, err := r.hhb.BuildIndexMap(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.HTML(http.StatusOK, "index.html", md)
	})
}

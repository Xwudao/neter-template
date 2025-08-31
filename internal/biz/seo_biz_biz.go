package biz

import (
	"bytes"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/Xwudao/neter-template/internal/domain/payloads"
	"github.com/Xwudao/neter-template/internal/system"
)

type SeoBizBiz struct {
	log    *zap.SugaredLogger
	appCtx *system.AppContext
}

func NewSeoBizBiz(log *zap.SugaredLogger, appCtx *system.AppContext) *SeoBizBiz {
	return &SeoBizBiz{
		log:    log.Named("seo-biz-biz"),
		appCtx: appCtx,
	}
}

func (h *SeoBizBiz) SEO(html []byte, c *gin.Context) (*payloads.SeoPayload, error) {
	var rtn = payloads.SeoPayload{
		Ret:        html,
		StatusCode: http.StatusOK,
	}

	var doc, err = goquery.NewDocumentFromReader(bytes.NewReader(html))
	if err != nil {
		return nil, err
	}

	var isDarkCookie, _ = c.Cookie("is_dark")

	if isDarkCookie == "true" {
		doc.Find("html").AddClass("dark")
		doc.Find("body").SetAttr("theme-mode", "dark")
	}

	ret, err := doc.Html()
	if err != nil {
		return nil, err
	}

	rtn.Ret = []byte(ret)

	return &rtn, nil

}

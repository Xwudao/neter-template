package biz

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/Xwudao/neter-template/internal/domain/models"
	"github.com/Xwudao/neter-template/internal/domain/payloads"
	"github.com/Xwudao/neter-template/internal/system"
	"github.com/Xwudao/neter-template/pkg/enum"
	"github.com/Xwudao/neter-template/pkg/utils/jwt"
)

type HtmlHelpBiz struct {
	log    *zap.SugaredLogger
	appCtx *system.AppContext

	scb *SiteConfigBiz
	dlb *DataListBiz
	jwt *jwt.Client
}

func NewHtmlHelpBiz(log *zap.SugaredLogger, appCtx *system.AppContext, scb *SiteConfigBiz, dlb *DataListBiz, jwt *jwt.Client) *HtmlHelpBiz {
	return &HtmlHelpBiz{
		log:    log.Named("html-help-biz"),
		appCtx: appCtx,
		scb:    scb,
		dlb:    dlb,
		jwt:    jwt,
	}
}
func (h *HtmlHelpBiz) BuildBase(c *gin.Context) (*payloads.HtmlBaseData, *models.SiteInfoConfig, error) {
	var cnf models.SiteInfoConfig
	if err := h.scb.GetConfig(enum.ConfigKeySiteInfo, &cnf); err != nil {
		return nil, nil, err
	}

	var (
		isDark, _ = c.Cookie("is_dark")
		token, _  = c.Cookie("token")
		logged    bool
	)

	if _, err := h.jwt.Parse(token); err == nil {
		logged = true
	}

	var m payloads.HtmlBaseData
	m.IsDark = isDark == "true"
	m.SiteName = cnf.SiteName
	m.SubTitle = cnf.SubTitle
	m.SiteDesc = cnf.SiteDesc
	m.SiteUrl = cnf.SiteUrl
	m.SiteImage = cnf.SiteImage // OpenGraph
	m.SiteLogo = cnf.SiteLogo
	m.SiteKeywords = cnf.SiteKeywords
	m.Disclaimer = cnf.Disclaimer
	m.SiteMetaScript = cnf.SitMetaScript

	m.Logged = logged

	return &m, &cnf, nil
}
func (h *HtmlHelpBiz) BuildIndexMap(c *gin.Context) (*payloads.IndexMap, error) {
	var base, cnf, err = h.BuildBase(c)
	if err != nil {
		return nil, err
	}

	var m payloads.IndexMap

	m.HtmlBaseData = *base
	m.Title = cnf.SiteTitle
	m.MainTitle = cnf.MainTitle

	/*meta*/
	var (
		canonical = m.SiteUrl
	)
	m.Canonical = canonical

	return &m, nil
}

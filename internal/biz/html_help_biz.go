package biz

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/Xwudao/neter-template/internal/domain/models"
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

func NewHtmlHelpBiz(
	log *zap.SugaredLogger,
	appCtx *system.AppContext,
	scb *SiteConfigBiz,
	dlb *DataListBiz,
	jwt *jwt.Client,
) *HtmlHelpBiz {
	return &HtmlHelpBiz{
		log:    log.Named("html-help-biz"),
		appCtx: appCtx,
		scb:    scb,
		dlb:    dlb,
		jwt:    jwt,
	}
}

func (h *HtmlHelpBiz) BuildBase(c *gin.Context) (*models.HtmlBaseModel, error) {
	var cnf models.SiteInfoConfig
	if err := h.scb.GetConfig(enum.ConfigKeySiteInfo, &cnf); err != nil {
		return nil, err
	}

	var (
		isDark, _ = c.Cookie("is_dark")
		token, _  = c.Cookie("token")
		logged    bool
	)

	if _, err := h.jwt.Parse(token); err == nil {
		logged = true
	}

	var m models.HtmlBaseModel
	m.IsDark = isDark == "true"
	m.Logged = logged

	m.SiteInfoConfig = cnf

	return &m, nil
}

func (h *HtmlHelpBiz) BuildIndexMap(c *gin.Context) (*models.IndexHtmlModel, error) {
	var base, err = h.BuildBase(c)
	if err != nil {
		return nil, err
	}

	var m models.IndexHtmlModel

	m.HtmlBaseModel = *base

	return &m, nil
}

package biz

import (
	"github.com/gin-gonic/gin"

	"github.com/Xwudao/neter-template/internal/domain/models"
)

// HtmlHelpBizIface is the interface consumed by route handlers.
type HtmlHelpBizIface interface {
	BuildBase(c *gin.Context) (*models.HtmlBaseModel, error)
	BuildIndexMap(c *gin.Context) (*models.IndexHtmlModel, error)
}

// Compile-time assertion: *HtmlHelpBiz must satisfy HtmlHelpBizIface.
var _ HtmlHelpBizIface = (*HtmlHelpBiz)(nil)

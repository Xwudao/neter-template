package params

import (
	"github.com/Xwudao/neter-template/internal/routes/valid"
)

type CreateSiteConfigParams struct {
	Name   string `json:"name" binding:"required"`
	Config string `json:"config" binding:"required"`
}

func (c *CreateSiteConfigParams) GetMessages() valid.ValidatorMessages {
	return valid.ValidatorMessages{
		"Name.required":   "名称不能为空",
		"Config.required": "配置不能为空",
	}
}

type UpdateSiteConfigParams struct {
	Name   string `json:"name" binding:"required"`
	Config string `json:"config" binding:"required"`
}

func (u *UpdateSiteConfigParams) GetMessages() valid.ValidatorMessages {
	return valid.ValidatorMessages{
		"Name.required":   "ID不能为空",
		"Config.required": "配置不能为空",
	}
}

type WriteFileParams struct {
	Filename string `json:"filename" binding:"required"`
	Data     string `json:"data" binding:"required"`
}

func (w *WriteFileParams) GetMessages() valid.ValidatorMessages {
	return valid.ValidatorMessages{
		"Filename.required": "文件名不能为空",
		"Data.required":     "数据不能为空",
	}
}

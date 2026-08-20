package params

import (
	"github.com/Xwudao/neter-template/internal/validate"
)

type CreateSiteConfigParams struct {
	Name   string `json:"name"`
	Config string `json:"config"`
}

func (c *CreateSiteConfigParams) Validate() error {
	return validate.Validate(
		validate.Field("name", c.Name,
			validate.Message("名称不能为空", validate.Required()),
		),
		validate.Field("config", c.Config,
			validate.Message("配置不能为空", validate.Required()),
		),
	)
}

type UpdateSiteConfigParams struct {
	Name   string `json:"name"`
	Config string `json:"config"`
}

func (u *UpdateSiteConfigParams) Validate() error {
	return validate.Validate(
		validate.Field("name", u.Name,
			validate.Message("ID不能为空", validate.Required()),
		),
		validate.Field("config", u.Config,
			validate.Message("配置不能为空", validate.Required()),
		),
	)
}

type WriteFileParams struct {
	Filename string `json:"filename"`
	Data     string `json:"data"`
}

func (w *WriteFileParams) Validate() error {
	return validate.Validate(
		validate.Field("filename", w.Filename,
			validate.Message("文件名不能为空", validate.Required()),
		),
		validate.Field("data", w.Data,
			validate.Message("数据不能为空", validate.Required()),
		),
	)
}

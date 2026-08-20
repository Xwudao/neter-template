package params

import (
	"mime/multipart"

	"github.com/Xwudao/neter-template/internal/validate"
)

type UploadToS3Params struct {
	Prefix string                `json:"prefix" form:"prefix"`
	Object string                `json:"object" form:"object"`
	File   *multipart.FileHeader `json:"file" form:"file"`
}

func (u *UploadToS3Params) Validate() error {
	return validate.Validate(
		validate.Field("prefix", u.Prefix,
			validate.Required(),
		),
		validate.Field("object", u.Object,
			validate.Required(),
		),
		validate.Field("file", u.File,
			validate.NotZero[*multipart.FileHeader](),
		),
	)
}

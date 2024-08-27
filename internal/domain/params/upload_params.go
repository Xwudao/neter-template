package params

import (
	"mime/multipart"
)

type UploadToS3Params struct {
	Prefix string `json:"prefix" binding:"required" form:"prefix"`
	Object string `json:"object" binding:"required" form:"object"`

	File *multipart.FileHeader `json:"file" binding:"required" form:"file"`
}

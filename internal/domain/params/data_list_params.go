package params

import (
	"go-kitboxpro/internal/routes/valid"
)

// CreateDataListParams 创建参数
type CreateDataListParams struct {
	Label string `json:"label" binding:"required"`
	Key   string `json:"key" binding:"required"`
	Kind  string `json:"kind" binding:"required"`
	Value string `json:"value" binding:"required"`

	ItemOrder int `json:"item_order"`
}

func (c *CreateDataListParams) Optimize() error {
	if c.ItemOrder == 0 {
		c.ItemOrder = 1
	}
	return nil
}

func (c *CreateDataListParams) GetMessages() valid.ValidatorMessages {
	return valid.ValidatorMessages{
		"Label.required": "标签必填",
		"Kind.required":  "分类必填",
		"Value.required": "内容必填",
		"Key.required":   "Key必填",
	}
}

// UpdateDataListParams 更新参数
type UpdateDataListParams struct {
	ID        int64  `json:"id" binding:"required"`
	Key       string `json:"key" binding:"required"`
	Value     string `json:"value" binding:"required"`
	ItemOrder *int   `json:"item_order"`
}

func (u *UpdateDataListParams) GetMessages() valid.ValidatorMessages {
	return valid.ValidatorMessages{
		"ID.required":    "ID必填",
		"Key.required":   "Key必填",
		"Value.required": "内容必填",
	}
}

type ListDataByKindParams struct {
	Kind string `json:"kind" binding:"required" form:"kind"`
	Page int    `json:"page" binding:"min=1" form:"page"`
	Size int    `json:"size" binding:"min=1,max=100" form:"size"`

	Offset int `json:"-" form:"-"`
}

func (l *ListDataByKindParams) Optimize() error {
	l.Offset = (l.Page - 1) * l.Size
	return nil
}

func (l *ListDataByKindParams) GetMessages() valid.ValidatorMessages {
	return valid.ValidatorMessages{
		"Kind.required": "Kind必填",
		"Page.min":      "Page最小值为1",
		"Size.min":      "Size最小值为1",
		"Size.max":      "Size最大值为100",
	}
}

// GetDataListSortDataParams 获取排序数据
type GetDataListSortDataParams struct {
	Kind string `json:"kind" binding:"required" form:"kind"`
}

func (g *GetDataListSortDataParams) GetMessages() valid.ValidatorMessages {
	return valid.ValidatorMessages{
		"Kind.required": "Kind必填",
	}
}

type GetAllDataListByKindsParams struct {
	Kinds   []string `json:"kinds"`
	ByOrder string   `json:"by_order"`
}

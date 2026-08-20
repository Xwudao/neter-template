package params

import (
	"github.com/Xwudao/neter-template/internal/validate"
)

// CreateDataListParams 创建参数
type CreateDataListParams struct {
	Label string `json:"label"`
	Key   string `json:"key"`
	Kind  string `json:"kind"`
	Value string `json:"value"`

	ItemOrder int `json:"item_order"`
}

func (c *CreateDataListParams) Optimize() error {
	if c.ItemOrder == 0 {
		c.ItemOrder = 1
	}
	return nil
}

func (c *CreateDataListParams) Validate() error {
	return validate.Validate(
		validate.Field("label", c.Label,
			validate.Message("标签必填", validate.Required()),
		),
		validate.Field("key", c.Key,
			validate.Message("Key必填", validate.Required()),
		),
		validate.Field("kind", c.Kind,
			validate.Message("分类必填", validate.Required()),
		),
		validate.Field("value", c.Value,
			validate.Message("内容必填", validate.Required()),
		),
	)
}

// UpdateDataListParams 更新参数
type UpdateDataListParams struct {
	ID        int64  `json:"id"`
	Key       string `json:"key"`
	Value     string `json:"value"`
	ItemOrder *int   `json:"item_order"`
}

func (u *UpdateDataListParams) Validate() error {
	return validate.Validate(
		validate.Field("id", u.ID,
			validate.Message("ID必填", validate.NotZero[int64]()),
		),
		validate.Field("key", u.Key,
			validate.Message("Key必填", validate.Required()),
		),
		validate.Field("value", u.Value,
			validate.Message("内容必填", validate.Required()),
		),
	)
}

type ListDataByKindParams struct {
	Kind string `json:"kind" form:"kind"`
	Page int    `json:"page" form:"page"`
	Size int    `json:"size" form:"size"`

	Offset int `json:"-" form:"-"`
}

func (l *ListDataByKindParams) Optimize() error {
	l.Offset = (l.Page - 1) * l.Size
	return nil
}

func (l *ListDataByKindParams) Validate() error {
	return validate.Validate(
		validate.Field("page", l.Page,
			validate.Message("Page最小值为1", validate.Min(1)),
		),
		validate.Field("size", l.Size,
			validate.Message("Size最小值为1", validate.Min(1)),
			validate.Message("Size最大值为100", validate.Max(100)),
		),
	)
}

// GetDataListSortDataParams 获取排序数据
type GetDataListSortDataParams struct {
	Kind string `json:"kind" form:"kind"`
}

func (g *GetDataListSortDataParams) Validate() error {
	return validate.Validate(
		validate.Field("kind", g.Kind,
			validate.Message("Kind必填", validate.Required()),
		),
	)
}

type GetAllDataListByKindsParams struct {
	Kinds   []string `json:"kinds"`
	ByOrder string   `json:"by_order"`
}

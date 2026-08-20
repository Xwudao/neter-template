package params

import (
	"errors"

	"github.com/Xwudao/neter-template/internal/validate"
)

type DeleteIDParams struct {
	ID int64 `json:"id"`
}

func (d *DeleteIDParams) Validate() error {
	return validate.Validate(
		validate.Field("id", d.ID,
			validate.Message("ID必填", validate.NotZero[int64]()),
		),
	)
}

type ItemOrderParams struct {
	IDs    []int64 `json:"ids"`
	Orders []int   `json:"orders"`
}

func (i *ItemOrderParams) Optimize() error {
	if len(i.IDs) != len(i.Orders) {
		return errors.New("ID和排序数量不一致")
	}
	return nil
}

func (i *ItemOrderParams) Validate() error {
	return validate.Validate(
		validate.Field("ids", i.IDs,
			validate.Message("ID必填", validate.MinItems[int64](1)),
		),
		validate.Field("orders", i.Orders,
			validate.Message("排序必填", validate.MinItems[int](1)),
		),
	)
}

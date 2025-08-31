package params

import (
	"errors"

	"github.com/Xwudao/neter-template/internal/routes/valid"
)

type DeleteIDParams struct {
	ID int64 `json:"id" binding:"required"`
}

func (d *DeleteIDParams) GetMessages() valid.ValidatorMessages {
	return valid.ValidatorMessages{
		"Name.required": "ID必填",
	}
}

type ItemOrderParams struct {
	IDs    []int64 `json:"ids" binding:"required"`
	Orders []int   `json:"orders" binding:"required"`
}

func (i *ItemOrderParams) Optimize() error {
	if len(i.IDs) != len(i.Orders) {
		return errors.New("ID和排序数量不一致")
	}
	return nil
}

func (i *ItemOrderParams) GetMessages() valid.ValidatorMessages {
	return valid.ValidatorMessages{
		"IDs.required":    "ID必填",
		"Orders.required": "排序必填",
	}
}

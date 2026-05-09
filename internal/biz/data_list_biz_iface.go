package biz

import (
	"context"

	"github.com/Xwudao/neter-template/internal/data/ent"
	"github.com/Xwudao/neter-template/internal/domain/params"
)

// DataListBizIface is the interface consumed by route handlers.
type DataListBizIface interface {
	Delete(ctx context.Context, id int64) error
	Get(ctx context.Context, id int64) (*ent.DataList, error)
	Create(ctx context.Context, p *params.CreateDataListParams) (*ent.DataList, error)
	Update(ctx context.Context, p *params.UpdateDataListParams) (*ent.DataList, error)
	GetAll(ctx context.Context) ([]*ent.DataList, error)
	ListByKind(ctx context.Context, p *params.ListDataByKindParams) ([]*ent.DataList, int, error)
	GetAllByKinds(ctx context.Context, p *params.GetAllDataListByKindsParams) ([]*ent.DataList, error)
	GetSortData(ctx context.Context, p *params.GetDataListSortDataParams) ([]*ent.DataList, error)
	UpdateOrder(ctx context.Context, p *params.ItemOrderParams) error
}

// Compile-time assertion: *DataListBiz must satisfy DataListBizIface.
var _ DataListBizIface = (*DataListBiz)(nil)

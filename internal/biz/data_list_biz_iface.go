package biz

import (
	"context"

	"github.com/Xwudao/neter-template/internal/data/sqlc"
	"github.com/Xwudao/neter-template/internal/domain/params"
)

// DataListBizIface is the interface consumed by route handlers.
type DataListBizIface interface {
	Delete(ctx context.Context, id int64) error
	Get(ctx context.Context, id int64) (*sqlc.DataList, error)
	Create(ctx context.Context, p *params.CreateDataListParams) (*sqlc.DataList, error)
	Update(ctx context.Context, p *params.UpdateDataListParams) (*sqlc.DataList, error)
	GetAll(ctx context.Context) ([]*sqlc.DataList, error)
	ListByKind(ctx context.Context, p *params.ListDataByKindParams) ([]*sqlc.DataList, int, error)
	GetAllByKinds(ctx context.Context, p *params.GetAllDataListByKindsParams) ([]*sqlc.DataList, error)
	GetSortData(ctx context.Context, p *params.GetDataListSortDataParams) ([]*sqlc.DataList, error)
	UpdateOrder(ctx context.Context, p *params.ItemOrderParams) error
}

// Compile-time assertion: *DataListBiz must satisfy DataListBizIface.
var _ DataListBizIface = (*DataListBiz)(nil)

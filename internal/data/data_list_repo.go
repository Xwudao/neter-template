package data

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"

	"github.com/Xwudao/neter-template/internal/biz"
	"github.com/Xwudao/neter-template/internal/data/sqlc"
	"github.com/Xwudao/neter-template/internal/domain/params"
	"github.com/Xwudao/neter-template/internal/system"
)

var _ biz.DataListRepository = (*dataListRepository)(nil)

type dataListRepository struct {
	appCtx *system.AppContext
	data   *Data
}

func NewDataListRepository(appCtx *system.AppContext, data *Data) biz.DataListRepository {
	return &dataListRepository{appCtx: appCtx, data: data}
}

func (u *dataListRepository) GetAll(ctx context.Context) ([]*sqlc.DataList, error) {
	return u.data.Queries.ListDataLists(ctx, "")
}

func (u *dataListRepository) DeleteByID(ctx context.Context, id int64) error {
	_, err := u.data.Queries.DeleteDataList(ctx, id)
	return err
}

func (u *dataListRepository) GetByID(ctx context.Context, id int64) (*sqlc.DataList, error) {
	return u.data.Queries.GetDataList(ctx, id)
}

func (u *dataListRepository) GetSortData(ctx context.Context, p *params.GetDataListSortDataParams) ([]*sqlc.DataList, error) {
	return u.data.Queries.ListDataListSortData(ctx, p.Kind)
}

func (u *dataListRepository) Create(ctx context.Context, p *params.CreateDataListParams) (*sqlc.DataList, error) {
	return u.data.Queries.CreateDataList(ctx, sqlc.CreateDataListParams{
		Label: p.Label, Kind: p.Kind, Key: p.Key, Value: p.Value, ItemOrder: int32(p.ItemOrder),
	})
}

func (u *dataListRepository) Update(ctx context.Context, p *params.UpdateDataListParams) (*sqlc.DataList, error) {
	itemOrder := pgtype.Int4{}
	if p.ItemOrder != nil {
		itemOrder = pgtype.Int4{Int32: int32(*p.ItemOrder), Valid: true}
	}
	return u.data.Queries.UpdateDataList(ctx, sqlc.UpdateDataListParams{
		ID: p.ID, Key: p.Key, Value: p.Value, ItemOrder: itemOrder,
	})
}

func (u *dataListRepository) GetAllByKinds(ctx context.Context, p *params.GetAllDataListByKindsParams) ([]*sqlc.DataList, error) {
	switch p.ByOrder {
	case "asc":
		return u.data.Queries.ListDataListsByKindsAsc(ctx, p.Kinds)
	case "desc":
		return u.data.Queries.ListDataListsByKindsDesc(ctx, p.Kinds)
	default:
		return u.data.Queries.ListDataListsByKinds(ctx, p.Kinds)
	}
}

func (u *dataListRepository) ListByKind(ctx context.Context, p *params.ListDataByKindParams) ([]*sqlc.DataList, int, error) {
	total, err := u.data.Queries.CountDataListsByKind(ctx, p.Kind)
	if err != nil {
		return nil, 0, err
	}
	items, err := u.data.Queries.ListDataListsByKindPage(ctx, sqlc.ListDataListsByKindPageParams{
		Kind: p.Kind, PageOffset: int32(p.Offset), PageSize: int32(p.Size),
	})
	return items, int(total), err
}

func (u *dataListRepository) UpdateOrder(ctx context.Context, p *params.ItemOrderParams) error {
	tx, err := u.data.Pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer func() { _ = tx.Rollback(ctx) }()

	queries := u.data.Queries.WithTx(tx)
	for i, id := range p.IDs {
		if _, err = queries.UpdateDataListOrder(ctx, sqlc.UpdateDataListOrderParams{ID: id, ItemOrder: int32(p.Orders[i])}); err != nil {
			return err
		}
	}
	return tx.Commit(ctx)
}

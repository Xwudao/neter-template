package data

import (
	"context"

	"entgo.io/ent/dialect/sql"

	"github.com/Xwudao/neter-template/internal/biz"
	"github.com/Xwudao/neter-template/internal/data/ent"
	"github.com/Xwudao/neter-template/internal/data/ent/datalist"
	"github.com/Xwudao/neter-template/internal/domain/params"
	"github.com/Xwudao/neter-template/internal/system"
)

var _ biz.DataListRepository = (*dataListRepository)(nil)

type dataListRepository struct {
	appCtx *system.AppContext
	data   *Data
}

func NewDataListRepository(appCtx *system.AppContext, data *Data) biz.DataListRepository {
	return &dataListRepository{
		appCtx: appCtx,
		data:   data,
	}
}

func (u *dataListRepository) GetAll(ctx context.Context) ([]*ent.DataList, error) {
	return u.data.Client.DataList.Query().All(ctx)
}

func (u *dataListRepository) DeleteByID(ctx context.Context, id int64) error {
	return u.data.Client.DataList.DeleteOneID(id).Exec(ctx)
}

func (u *dataListRepository) GetByID(ctx context.Context, id int64) (*ent.DataList, error) {
	return u.data.Client.DataList.Get(ctx, id)
}

// GetSortData 获取排序的字段id,name
func (u *dataListRepository) GetSortData(ctx context.Context, p *params.GetDataListSortDataParams) ([]*ent.DataList, error) {
	var builder = u.data.Client.DataList.Query().Where(datalist.KindEQ(p.Kind)).
		Order(datalist.ByItemOrder(sql.OrderDesc()))

	return builder.All(ctx)
}

func (u *dataListRepository) Create(ctx context.Context, p *params.CreateDataListParams) (*ent.DataList, error) {
	return u.data.Client.DataList.Create().SetItemOrder(p.ItemOrder).SetKey(p.Key).SetKind(p.Kind).
		SetValue(p.Value).SetLabel(p.Label).
		Save(ctx)
}

func (u *dataListRepository) Update(ctx context.Context, p *params.UpdateDataListParams) (*ent.DataList, error) {
	return u.data.Client.DataList.UpdateOneID(p.ID).SetKey(p.Key).SetNillableItemOrder(p.ItemOrder).
		SetValue(p.Value).Save(ctx)
}

// GetAllByKinds 根据Kind查询
func (u *dataListRepository) GetAllByKinds(ctx context.Context, p *params.GetAllDataListByKindsParams) ([]*ent.DataList, error) {
	var builder = u.data.Client.DataList.Query().Where(datalist.KindIn(p.Kinds...))

	switch p.ByOrder {
	case "asc":
		builder.Order(datalist.ByItemOrder())
	case "desc":
		builder.Order(datalist.ByItemOrder(sql.OrderDesc()))
	}

	return builder.All(ctx)
}

// ListByKind 根据Kind查询
func (u *dataListRepository) ListByKind(ctx context.Context, p *params.ListDataByKindParams) ([]*ent.DataList, int, error) {
	var builder = u.data.Client.DataList.Query()

	if p.Kind != "" {
		builder.Where(datalist.KindEQ(p.Kind))
	}

	var total = builder.CountX(ctx)

	rtn, err := builder.Offset(p.Offset).Limit(p.Size).All(ctx)
	return rtn, total, err
}

func (u *dataListRepository) UpdateOrder(ctx context.Context, p *params.ItemOrderParams) error {
	for i, id := range p.IDs {
		if _, err := u.data.Client.DataList.UpdateOneID(id).SetItemOrder(p.Orders[i]).Save(ctx); err != nil {
			return err
		}
	}

	return nil
}

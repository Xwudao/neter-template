package biz

import (
	"context"

	"go.uber.org/zap"

	"github.com/Xwudao/neter-template/internal/data/ent"
	"github.com/Xwudao/neter-template/internal/domain/params"
	"github.com/Xwudao/neter-template/internal/system"
)

type DataListRepository interface {
	GetAll(ctx context.Context) ([]*ent.DataList, error)
	DeleteByID(ctx context.Context, id int64) error
	GetByID(ctx context.Context, id int64) (*ent.DataList, error)
	GetSortData(ctx context.Context, p *params.GetDataListSortDataParams) ([]*ent.DataList, error)
	Create(ctx context.Context, p *params.CreateDataListParams) (*ent.DataList, error)
	Update(ctx context.Context, p *params.UpdateDataListParams) (*ent.DataList, error)
	GetAllByKinds(ctx context.Context, kinds *params.GetAllDataListByKindsParams) ([]*ent.DataList, error)
	ListByKind(ctx context.Context, p *params.ListDataByKindParams) ([]*ent.DataList, int, error)
	UpdateOrder(ctx context.Context, p *params.ItemOrderParams) error
}

type DataListBiz struct {
	log    *zap.SugaredLogger
	appCtx *system.AppContext
	dlr    DataListRepository
}

func NewDataListBiz(log *zap.SugaredLogger, dlr DataListRepository, appCtx *system.AppContext) *DataListBiz {
	return &DataListBiz{
		log:    log.Named("data-list-biz"),
		appCtx: appCtx,
		dlr:    dlr,
	}
}

func (h *DataListBiz) Delete(ctx context.Context, id int64) error {
	return h.dlr.DeleteByID(ctx, id)
}

func (h *DataListBiz) Get(ctx context.Context, id int64) (*ent.DataList, error) {
	return h.dlr.GetByID(ctx, id)
}

func (h *DataListBiz) Create(ctx context.Context, p *params.CreateDataListParams) (*ent.DataList, error) {
	if err := p.Optimize(); err != nil {
		return nil, err
	}
	return h.dlr.Create(ctx, p)
}

func (h *DataListBiz) Update(ctx context.Context, p *params.UpdateDataListParams) (*ent.DataList, error) {
	return h.dlr.Update(ctx, p)
}

func (h *DataListBiz) GetAll(ctx context.Context) ([]*ent.DataList, error) {
	return h.dlr.GetAll(ctx)
}

func (h *DataListBiz) ListByKind(ctx context.Context, p *params.ListDataByKindParams) ([]*ent.DataList, int, error) {
	if err := p.Optimize(); err != nil {
		return nil, 0, err
	}
	return h.dlr.ListByKind(ctx, p)
}

func (h *DataListBiz) GetAllByKinds(ctx context.Context, p *params.GetAllDataListByKindsParams) ([]*ent.DataList, error) {
	return h.dlr.GetAllByKinds(ctx, p)
}

func (h *DataListBiz) GetSortData(ctx context.Context, p *params.GetDataListSortDataParams) ([]*ent.DataList, error) {
	return h.dlr.GetSortData(ctx, p)
}

func (h *DataListBiz) UpdateOrder(ctx context.Context, p *params.ItemOrderParams) error {
	if err := p.Optimize(); err != nil {
		return err
	}
	return h.dlr.UpdateOrder(ctx, p)
}

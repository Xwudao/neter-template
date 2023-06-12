package biz

import (
	"context"

	"go.uber.org/zap"

	"github.com/Xwudao/neter-template/internal/data/ent"
	"github.com/Xwudao/neter-template/internal/system"
)

type UserRepository interface {
	GetAll(ctx context.Context) ([]*ent.User, error)
	DeleteByID(ctx context.Context, id int64) error
	GetByID(ctx context.Context, id int64) (*ent.User, error)
	Create(ctx context.Context) (*ent.User, error)
}

type UserBiz struct {
	log    *zap.SugaredLogger
	appCtx *system.AppContext
	ur     UserRepository
}

func NewUserBiz(log *zap.SugaredLogger /*ur UserRepository,*/, appCtx *system.AppContext) *UserBiz {
	return &UserBiz{
		log:    log.Named("user-biz"),
		appCtx: appCtx,
		//ur:  ur,
	}
}

func (h *UserBiz) Index() string {
	panic("TODO implement")
}

func (h *UserBiz) Delete(ctx context.Context, id int64) error {
	return h.ur.DeleteByID(ctx, id)
}

func (h *UserBiz) Get(ctx context.Context, id int64) (*ent.User, error) {
	return h.ur.GetByID(ctx, id)
}

func (h *UserBiz) Create(ctx context.Context) (*ent.User, error) {
	return h.ur.Create(ctx)
}

func (h *UserBiz) GetAll(ctx context.Context) ([]*ent.User, error) {
	return h.ur.GetAll(ctx)
}

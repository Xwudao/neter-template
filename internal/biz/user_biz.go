package biz

import (
	"context"

	"go.uber.org/zap"

	"github.com/Xwudao/neter-template/internal/data/ent"
	"github.com/Xwudao/neter-template/internal/system"
)

type UserRepository interface {
	GetAll() ([]*ent.User, error)
	DeleteByID(id int64) error
	GetByID(id int64) (*ent.User, error)
	Create() (*ent.User, error)
}

type UserBiz struct {
	log *zap.SugaredLogger
	ctx context.Context
	ur  UserRepository
}

func NewUserBiz(log *zap.SugaredLogger, ur UserRepository, appCtx *system.AppContext) *UserBiz {
	return &UserBiz{
		log: log.Named("user-biz"),
		ctx: appCtx.Ctx,
		ur:  ur,
	}
}

func (h *UserBiz) Index() string {
	panic("TODO implement")
}

func (h *UserBiz) Delete(id int64) error {
	return h.ur.DeleteByID(id)
}

func (h *UserBiz) Get(id int64) (*ent.User, error) {
	return h.ur.GetByID(id)
}

func (h *UserBiz) Create() (*ent.User, error) {
	return h.ur.Create()
}

func (h *UserBiz) GetAll() ([]*ent.User, error) {
	return h.ur.GetAll()
}

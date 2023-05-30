package data

import (
	"context"

	"github.com/Xwudao/neter-template/internal/biz"
	"github.com/Xwudao/neter-template/internal/data/ent"
	"github.com/Xwudao/neter-template/internal/system"
)

var _ biz.UserRepository = (*userRepository)(nil)

type userRepository struct {
	ctx  context.Context
	data *Data
}

func NewUserRepository(appCtx *system.AppContext, data *Data) biz.UserRepository {
	return &userRepository{
		ctx:  appCtx.Ctx,
		data: data,
	}
}

func (u *userRepository) GetAll() ([]*ent.User, error) {
	return u.data.Client.User.Query().All(u.ctx)
}

func (u *userRepository) DeleteByID(id int64) error {
	return u.data.Client.User.DeleteOneID(id).Exec(u.ctx)
}

func (u *userRepository) GetByID(id int64) (*ent.User, error) {
	return u.data.Client.User.Get(u.ctx, id)
}

func (u *userRepository) Create() (*ent.User, error) {
	// todo add set fields
	return u.data.Client.User.Create().Save(u.ctx)
}

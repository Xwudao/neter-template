package data

import (
	"context"

	"github.com/Xwudao/neter-template/internal/biz"
	"github.com/Xwudao/neter-template/internal/data/ent"
	"github.com/Xwudao/neter-template/internal/system"
)

var _ biz.UserRepository = (*userRepository)(nil)

type userRepository struct {
	appCtx *system.AppContext
	data   *Data
}

func NewUserRepository(appCtx *system.AppContext, data *Data) biz.UserRepository {
	return &userRepository{
		appCtx: appCtx,
		data:   data,
	}
}

func (u *userRepository) GetAll(ctx context.Context) ([]*ent.User, error) {
	return u.data.Client.User.Query().All(ctx)
}

func (u *userRepository) DeleteByID(ctx context.Context, id int64) error {
	return u.data.Client.User.DeleteOneID(id).Exec(ctx)
}

func (u *userRepository) GetByID(ctx context.Context, id int64) (*ent.User, error) {
	return u.data.Client.User.Get(ctx, id)
}

func (u *userRepository) Create(ctx context.Context) (*ent.User, error) {
	// todo add set fields
	return u.data.Client.User.Create().Save(ctx)
}

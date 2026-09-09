package data

import (
	"context"

	"github.com/Xwudao/neter-template/internal/biz"
	"github.com/Xwudao/neter-template/internal/data/sqlc"
	"github.com/Xwudao/neter-template/internal/domain/params"
	"github.com/Xwudao/neter-template/internal/system"
)

var _ biz.UserRepository = (*userRepository)(nil)

type userRepository struct {
	appCtx *system.AppContext
	data   *Data
}

func NewUserRepository(appCtx *system.AppContext, data *Data) biz.UserRepository {
	return &userRepository{appCtx: appCtx, data: data}
}

func (u *userRepository) GetAll(ctx context.Context) ([]*sqlc.User, error) {
	return u.data.Queries.ListUsers(ctx)
}

func (u *userRepository) DeleteByID(ctx context.Context, id int64) error {
	_, err := u.data.Queries.DeleteUser(ctx, id)
	return err
}

func (u *userRepository) GetByID(ctx context.Context, id int64) (*sqlc.User, error) {
	return u.data.Queries.GetUser(ctx, id)
}

func (u *userRepository) GetBy(ctx context.Context, p *params.GetUserByParams) (*sqlc.User, error) {
	if p.Username != "" {
		return u.data.Queries.GetUserByUsername(ctx, p.Username)
	}
	return u.data.Queries.GetUser(ctx, p.ID)
}

func (u *userRepository) Create(ctx context.Context, p *params.CreateUserParams) (*sqlc.User, error) {
	return u.data.Queries.CreateUser(ctx, sqlc.CreateUserParams{
		Username: p.Username,
		Password: p.Password,
		Role:     p.Role,
	})
}

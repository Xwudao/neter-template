package data

import (
	"cmp"
	"context"

	"github.com/Xwudao/neter-template/internal/biz"
	"github.com/Xwudao/neter-template/internal/data/ent"
	"github.com/Xwudao/neter-template/internal/data/ent/user"
	"github.com/Xwudao/neter-template/internal/domain/params"
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

func (u *userRepository) GetBy(ctx context.Context, p *params.GetUserByParams) (*ent.User, error) {
	var builder = u.data.Client.User.Query()
	if p.Username != "" {
		builder.Where(user.UsernameEQ(p.Username))
	}

	if p.ID != 0 {
		builder.Where(user.IDEQ(p.ID))
	}

	return builder.First(ctx)
}

func (u *userRepository) Create(ctx context.Context, p *params.CreateUserParams) (*ent.User, error) {
	return u.data.Client.User.Create().SetUsername(p.Username).SetPassword(p.Password).
		SetRole(cmp.Or(p.Role, user.RoleUser)).Save(ctx)
}

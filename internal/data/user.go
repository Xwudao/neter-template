package data

import (
	"github.com/Xwudao/neter-template/internal/core"
)

type UserRepository struct {
	data *Data
	ctx  *core.AppContext
}

func NewUserRepository(ctx *core.AppContext, data *Data) *UserRepository {
	return &UserRepository{data: data, ctx: ctx}
}

func (ur *UserRepository) CreateUser(name string) error {
	return ur.data.client.User.Create().SetName(name).Exec(ur.ctx.Ctx)
}

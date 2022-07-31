package data

import (
	"github.com/Xwudao/neter-template/internal/biz"
	"github.com/Xwudao/neter-template/internal/core"
)

type userRepository struct {
	data *Data
	ctx  *core.AppContext
}

func NewUserRepository(ctx *core.AppContext, data *Data) biz.UserRepository {
	return &userRepository{data: data, ctx: ctx}
}

func (ur *userRepository) CreateUser(name string) error {
	return ur.data.client.User.Create().SetName(name).Exec(ur.ctx.Ctx)
}

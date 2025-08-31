package params

import (
	"github.com/Xwudao/neter-template/internal/data/ent/user"
	"github.com/Xwudao/neter-template/internal/routes/valid"
)

type CreateUserParams struct {
	Username string    `json:"username" binding:"required"`
	Password string    `json:"password" binding:"required"`
	Role     user.Role `json:"-"`
}

func (c *CreateUserParams) GetMessages() valid.ValidatorMessages {
	return valid.ValidatorMessages{
		"Username.required": "用户名不能为空",
		"Password.required": "密码不能为空",
	}
}

type GetUserByParams struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}

type UserLoginParams struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *UserLoginParams) GetMessages() valid.ValidatorMessages {
	return valid.ValidatorMessages{
		"Username.required": "用户名不能为空",
		"Password.required": "密码不能为空",
	}
}

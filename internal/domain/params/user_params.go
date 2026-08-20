package params

import (
	"github.com/Xwudao/neter-template/internal/data/ent/user"
	"github.com/Xwudao/neter-template/internal/validate"
)

type CreateUserParams struct {
	Username string    `json:"username"`
	Password string    `json:"password"`
	Role     user.Role `json:"-"`
}

func (c *CreateUserParams) Validate() error {
	return validate.Validate(
		validate.Field("username", c.Username,
			validate.Message("用户名不能为空", validate.Required()),
		),
		validate.Field("password", c.Password,
			validate.Message("密码不能为空", validate.Required()),
		),
	)
}

type GetUserByParams struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}

type UserLoginParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *UserLoginParams) Validate() error {
	return validate.Validate(
		validate.Field("username", u.Username,
			validate.Message("用户名不能为空", validate.Required()),
		),
		validate.Field("password", u.Password,
			validate.Message("密码不能为空", validate.Required()),
		),
	)
}

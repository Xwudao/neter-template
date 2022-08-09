package data

import (
	"github.com/Xwudao/neter-template/internal/biz"
)

var _ biz.UserRepository = (*userRepository)(nil)

type userRepository struct {
}

func (u *userRepository) TodoFunc() error {
	//TODO implement me
	panic("implement me")
}
func NewUserRepository() biz.UserRepository {
	return &userRepository{}
}

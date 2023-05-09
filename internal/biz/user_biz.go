package biz

type UserRepository interface {
	TodoFunc() error
}

type UserBiz struct {
}

func NewUserBiz() *UserBiz {
	return &UserBiz{}
}

func (h *UserBiz) Index() string {
	panic("TODO implement")
}

package biz

type UserBiz struct {
}

func NewUserBiz() *UserBiz {
	return &UserBiz{}
}

type UserRepository interface {
	TodoFunc() error
}

func (h *UserBiz) Index() string {
	panic("TODO implement")
}

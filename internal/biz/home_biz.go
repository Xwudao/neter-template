package biz

import "fmt"

type HomeBiz struct {
}

func NewHomeBiz() *HomeBiz {
	return &HomeBiz{}
}

type UserRepository interface {
	CreateUser(name string) error
}

func (h *HomeBiz) SayHello(name string) string {
	return fmt.Sprintf("Hello %s", name)
}

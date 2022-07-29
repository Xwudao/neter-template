package biz

import "fmt"

type HomeBiz struct {
}

func NewHomeBiz() *HomeBiz {
	return &HomeBiz{}
}

func (h *HomeBiz) SayHello(name string) string {
	return fmt.Sprintf("Hello %s", name)
}

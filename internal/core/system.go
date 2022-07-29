package core

import (
	"github.com/Xwudao/neter-template/pkg/utils/bcrypt"
)

type InitSystem struct {
}

// InitConfig init some config in some package
func (i *InitSystem) InitConfig() {
	bcrypt.Init(bcrypt.WithCost(10))
}

func NewInitSystem() *InitSystem {
	i := &InitSystem{}
	i.InitConfig()

	return i
}

package system

import (
	"errors"

	"github.com/knadh/koanf/v2"

	"go-kitboxpro/pkg/utils/bcrypt"
)

type InitSystem struct {
	conf *koanf.Koanf
}

func NewInitSystem(conf *koanf.Koanf) *InitSystem {
	i := &InitSystem{
		conf: conf,
	}
	i.InitConfig()

	return i
}

// InitConfig init some config in some package
func (i *InitSystem) InitConfig() {
	bcrypt.Init(bcrypt.WithCost(10))
}

// CheckSystem check system
func (i *InitSystem) CheckSystem() error {
	jwtSecret := i.conf.String("jwt.secret")
	if len(jwtSecret) <= 12 {
		return errors.New("jwt secret is too short")
	}
	return nil
}

package system

import (
	"github.com/Xwudao/neter-template/pkg/utils/bcrypt"
	"github.com/Xwudao/neter-template/pkg/utils/cron"
)

type InitSystem struct {
	cron *cron.Cron
}

// InitConfig init some config in some package
func (i *InitSystem) InitConfig() {
	bcrypt.Init(bcrypt.WithCost(10))
}

func NewInitSystem(cron *cron.Cron) *InitSystem {
	i := &InitSystem{
		cron: cron,
	}
	i.InitConfig()
	i.cron.InitCron()

	return i
}

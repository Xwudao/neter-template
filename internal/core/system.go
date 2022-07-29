package core

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

//InitCron init cron job
func (i *InitSystem) InitCron() {
	// add cron jobs in here...
}

func NewInitSystem(cron *cron.Cron) *InitSystem {
	i := &InitSystem{
		cron: cron,
	}
	i.InitConfig()
	i.InitCron()

	return i
}

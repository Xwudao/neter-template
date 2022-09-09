package cron

import (
	"github.com/google/wire"
)

var ProviderCronSet = wire.NewSet(NewCron)

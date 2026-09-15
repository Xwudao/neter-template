package cron

import "github.com/Xwudao/loom"

var ProviderCronSet = loom.Module(loom.Provide(NewCron))

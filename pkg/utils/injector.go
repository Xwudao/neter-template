package utils

import (
	"github.com/google/wire"

	"github.com/Xwudao/neter-template/pkg/utils/cron"
)

var ProvideUtilSet = wire.NewSet(cron.NewCron)

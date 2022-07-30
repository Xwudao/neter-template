package utils

import (
	"github.com/google/wire"

	"github.com/Xwudao/neter-template/pkg/utils/cron"
	"github.com/Xwudao/neter-template/pkg/utils/jwt"
)

var (
	ProvideUtilSet = wire.NewSet(cron.NewCron, jwt.NewClient)
)

package utils

import (
	"github.com/Xwudao/loom"

	"github.com/Xwudao/neter-template/pkg/utils/jwt"
)

var ProvideUtilSet = loom.Module(loom.Provide(jwt.NewClient))

package utils

import (
	"github.com/google/wire"

	"github.com/Xwudao/neter-template/pkg/utils/jwt"
)

var (
	ProvideUtilSet = wire.NewSet(jwt.NewClient)
)

package utils

import (
	"github.com/google/wire"

	"go-kitboxpro/pkg/utils/jwt"
)

var (
	ProvideUtilSet = wire.NewSet(jwt.NewClient)
)

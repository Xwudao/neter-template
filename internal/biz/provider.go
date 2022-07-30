package biz

import "github.com/google/wire"

var ProviderBizSet = wire.NewSet(NewHomeBiz)

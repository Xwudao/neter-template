package config

import (
	"github.com/google/wire"
)

var ProviderConfigSet = wire.NewSet(
	NewKoanf,
	NewJwtConfigData,
	NewDBConfig,
	NewCorsConfig,
	NewProxyConfig,
	NewS3Config,
)

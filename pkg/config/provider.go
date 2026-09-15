package config

import "github.com/Xwudao/loom"

var ProviderConfigSet = loom.Module(
	loom.Provide(NewKoanf),
	loom.Provide(NewJwtConfigData),
	loom.Provide(NewDBConfig),
	loom.Provide(NewCorsConfig),
	loom.Provide(NewProxyConfig),
	loom.Provide(NewS3Config),
)

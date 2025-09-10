package payloads

import (
	"time"
)

type JwtConfig struct {
	Secret string        `json:"secret" koanf:"secret"`
	Expire time.Duration `json:"expire" koanf:"expire"`
	Issuer string        `json:"issuer" koanf:"issuer"`
}

type DBConfig struct {
	Host        string `json:"host" koanf:"host"`
	Port        int    `json:"port" koanf:"port"`
	Dialect     string `json:"dialect" koanf:"dialect"`
	Username    string `json:"username" koanf:"username"`
	Password    string `json:"password" koanf:"password"`
	Database    string `json:"database" koanf:"database"`
	AutoMigrate bool   `json:"autoMigrate" koanf:"autoMigrate"`
}

type CorsConfig struct {
	AllowOrigin      []string      `json:"allowOrigin" koanf:"allowOrigin"`
	AllowCredentials bool          `json:"allowCredentials" koanf:"allowCredentials"`
	MaxAge           time.Duration `json:"maxAge" koanf:"maxAge"`
}

type ProxyConfig struct {
	Username string `json:"username" koanf:"username"`
	Password string `json:"password" koanf:"password"`
	Addr     string `json:"addr" koanf:"addr"`
}

type S3Config struct {
	AccessKey string `json:"accessKey" koanf:"accessKey"`
	SecretKey string `json:"secretKey" koanf:"secret"`
	Endpoint  string `json:"endpoint" koanf:"endpoint"`
	Bucket    string `json:"bucket" koanf:"bucket"`
	ProxyUrl  string `json:"proxyUrl" koanf:"proxyUrl"`
}

type RedisConfig struct {
	Addr     string `json:"addr" koanf:"addr"`
	DB       int    `json:"db" koanf:"db"`
	Password string `json:"password" koanf:"password"`
	Prefix   string `json:"prefix" koanf:"prefix"`
}

func (r *RedisConfig) FormatKey(key string) string {
	if r.Prefix == "" {
		return key
	}

	return r.Prefix + ":" + key
}

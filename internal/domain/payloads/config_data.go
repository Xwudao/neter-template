package payloads

import (
	"time"
)

type JwtConfig struct {
	Secret string        `json:"secret" yaml:"secret"`
	Expire time.Duration `json:"expire" yaml:"expire"`
	Issuer string        `json:"issuer" yaml:"issuer"`
}

type DBConfig struct {
	Host        string `json:"host" yaml:"host"`
	Port        int    `json:"port" yaml:"port"`
	Dialect     string `json:"dialect" yaml:"dialect"`
	Username    string `json:"username" yaml:"username"`
	Password    string `json:"password" yaml:"password"`
	Database    string `json:"database" yaml:"database"`
	AutoMigrate bool   `json:"autoMigrate" yaml:"autoMigrate"`
}

type CorsConfig struct {
	AllowOrigin      []string      `json:"allowOrigin" yaml:"allowOrigin"`
	AllowCredentials bool          `json:"allowCredentials" yaml:"allowCredentials"`
	MaxAge           time.Duration `json:"maxAge" yaml:"maxAge"`
}

type ProxyConfig struct {
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
	Addr     string `json:"addr" yaml:"addr"`
}

type S3Config struct {
	AccessKey string `json:"accessKey" yaml:"accessKey"`
	SecretKey string `json:"secretKey" yaml:"secret"`
	Endpoint  string `json:"endpoint" yaml:"endpoint"`
	Bucket    string `json:"bucket" yaml:"bucket"`
	ProxyUrl  string `json:"proxyUrl" yaml:"proxyUrl"`
}

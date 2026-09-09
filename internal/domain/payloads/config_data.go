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
	Host     string `json:"host" koanf:"host"`
	Port     int    `json:"port" koanf:"port"`
	Username string `json:"username" koanf:"username"`
	Password string `json:"password" koanf:"password"`
	Database string `json:"database" koanf:"database"`

	// AutoMigrate applies pending migrations when the binary starts.
	// Override with NETER_AUTO_MIGRATE=true|false.
	AutoMigrate bool `json:"autoMigrate" koanf:"autoMigrate"`
	// MigrateDsn overrides the application DSN for migrations, e.g. a
	// dedicated DDL user. Override with NETER_MIGRATE_DSN.
	MigrateDsn string `json:"migrateDsn" koanf:"migrateDsn"`
	// MigratePath, when set, reads migrations from this directory on disk
	// instead of the embedded copy (development only).
	MigratePath string `json:"migratePath" koanf:"migratePath"`
}

type CorsConfig	 struct {
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

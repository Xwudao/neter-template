package config

import (
	"github.com/knadh/koanf/v2"

	"github.com/Xwudao/neter-template/internal/domain/payloads"
)

func NewJwtConfigData(conf *koanf.Koanf) (*payloads.JwtConfig, error) {
	var jwtConfig payloads.JwtConfig
	if err := conf.Unmarshal("jwt", &jwtConfig); err != nil {
		return nil, err
	}

	return &jwtConfig, nil
}

func NewDBConfig(conf *koanf.Koanf) (*payloads.DBConfig, error) {
	var dbConfig payloads.DBConfig
	if err := conf.Unmarshal("db", &dbConfig); err != nil {
		return nil, nil
	}

	return &dbConfig, nil
}

// NewCorsConfig reads the cors configuration from the koanf config
func NewCorsConfig(conf *koanf.Koanf) (*payloads.CorsConfig, error) {
	var corsConfig payloads.CorsConfig
	if err := conf.Unmarshal("cors", &corsConfig); err != nil {
		return nil, err
	}

	return &corsConfig, nil
}

// NewProxyConfig reads the proxy configuration from the koanf config
func NewProxyConfig(conf *koanf.Koanf) (*payloads.ProxyConfig, error) {
	var proxyConfig payloads.ProxyConfig
	if err := conf.Unmarshal("proxy", &proxyConfig); err != nil {
		return nil, err
	}

	return &proxyConfig, nil
}

func NewS3Config(conf *koanf.Koanf) (*payloads.S3Config, error) {
	var s3Config payloads.S3Config
	if err := conf.Unmarshal("s3", &s3Config); err != nil {
		return nil, err
	}

	return &s3Config, nil
}

package utils

import (
	"net/url"

	"golang.org/x/net/publicsuffix"
)

func ExtractDomain(rawURL string) (string, error) {
	res, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	return res.Hostname(), nil
}

// ExtractRootDomain 从 URL 中提取注册域（根域名）
func ExtractRootDomain(rawURL string) (string, error) {
	// 解析 URL
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	// 获取 Hostname（去掉端口）
	host := parsed.Hostname()

	// 利用 publicsuffix 提取注册域
	eTLDPlusOne, err := publicsuffix.EffectiveTLDPlusOne(host)
	if err != nil {
		return "", err
	}

	return eTLDPlusOne, nil
}

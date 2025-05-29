package utils

import (
	"net/url"
	"strings"

	"go-kitboxpro/internal/domain/payloads"
)

func BuildProxyUrl(proxyData *payloads.ProxyConfig) string {
	if proxyData == nil {
		return ""
	}

	parse, err := url.Parse(proxyData.Addr)
	if err != nil {
		return ""
	}
	var (
		port     = parse.Port()
		scheme   = parse.Scheme
		host     = parse.Hostname()
		username = proxyData.Username
		password = proxyData.Password
	)

	if username != "" && password != "" {
		return scheme + "://" + username + ":" + password + "@" + host + ":" + port
	}
	return scheme + "://" + host + ":" + port
}

// CleanUrl 去除url中的参数，只保留域名和路径
func CleanUrl(siteUrl string) (string, error) {
	parse, err := url.Parse(siteUrl)
	if err != nil {
		return "", err
	}
	var aimUrl = parse.Scheme + "://" + parse.Host + parse.Path
	return aimUrl, nil
}

// ParseFavicon 解析站点图标
func ParseFavicon(siteUrl, faviconPath string) (string, error) {
	if strings.HasPrefix(faviconPath, "http") {
		return faviconPath, nil
	}

	if faviconPath == "" {
		return "", nil
	}

	parse, err := url.Parse(siteUrl)
	if err != nil {
		return "", err
	}

	if faviconPath[0] == '/' {
		parse.Path = "" // 清空path，只保留域名
	}

	return parse.JoinPath(faviconPath).String(), nil
}

// ParseHostname 从url中解析出域名
func ParseHostname(siteUrl string) (string, error) {
	parse, err := url.Parse(siteUrl)
	if err != nil {
		return "", err
	}
	return parse.Hostname(), nil
}

// MustHostname 从url中解析出域名
func MustHostname(siteUrl string) string {
	hostname, err := ParseHostname(siteUrl)
	if err != nil {
		return ""
	}
	return hostname
}

package utils

import (
	"encoding/base64"
	"net/url"
)

// B64Encode encodes a string to base64
func B64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func AddUTM(siteUrl string) string {
	//	utm_source=v2fd.com
	rtn, err := ModifyURLParams(siteUrl, map[string]string{
		"utm_source": "v2fd.com",
	})
	if err != nil {
		return siteUrl
	}
	return rtn
}

func ModifyURLParams(inputURL string, params map[string]string) (string, error) {
	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		return "", err
	}

	query := parsedURL.Query()

	for key, value := range params {
		query.Set(key, value)
	}

	parsedURL.RawQuery = query.Encode()

	return parsedURL.String(), nil
}

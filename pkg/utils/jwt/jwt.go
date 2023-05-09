package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/knadh/koanf"
)

type Client struct {
	conf *koanf.Koanf
}

func NewClient(conf *koanf.Koanf) *Client {
	return &Client{
		conf: conf,
	}
}

func (c *Client) Generate(userID int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   userID,
		"sign_date": time.Now().Unix(),
	})

	signedString, err := token.SignedString([]byte(c.conf.MustString("jwt.secret")))
	if err != nil {
		return "", err
	}
	return signedString, nil
}

// Parse the token string
func (c *Client) Parse(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(c.conf.MustString("jwt.secret")), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

// Validate validate
func (c *Client) Validate(tokenString string) error {
	_, err := c.Parse(tokenString)
	return err
}

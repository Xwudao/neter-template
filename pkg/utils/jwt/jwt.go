package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"go-kitboxpro/internal/domain/payloads"
)

type CustomClaims struct {
	UserID int64 `json:"user_id"`
	jwt.RegisteredClaims
}

type Client struct {
	conf *payloads.JwtConfig
}

func NewClient(conf *payloads.JwtConfig) *Client {
	return &Client{
		conf: conf,
	}
}

func (c *Client) Generate(userID int64) (string, error) {
	//var d = c.conf.MustDuration("jwt.expire")

	var exp = time.Now().Add(c.conf.Expire)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    c.conf.Issuer,
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	})

	signedString, err := token.SignedString([]byte(c.conf.Secret))
	if err != nil {
		return "", err
	}
	return signedString, nil
}

// Parse the token string
func (c *Client) Parse(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(c.conf.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return jwt.MapClaims{
			"user_id": claims.UserID,
		}, nil
	}
	return nil, errors.New("invalid token")
}

// Validate validate
func (c *Client) Validate(tokenString string) error {
	_, err := c.Parse(tokenString)
	return err
}

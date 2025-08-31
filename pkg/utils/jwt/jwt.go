package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/Xwudao/neter-template/internal/domain/payloads"
)

var (
	ErrInvalidToken     = errors.New("invalid token")
	ErrExpiredToken     = errors.New("token has expired")
	ErrInvalidSignature = errors.New("invalid token signature")
	ErrTokenNotFound    = errors.New("token not found")
)

type CustomClaims struct {
	UserID int64 `json:"user_id"`
	jwt.RegisteredClaims
}

type Client struct {
	conf *payloads.JwtConfig
}

type TokenPair struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token,omitempty"`
	ExpiresAt    time.Time `json:"expires_at"`
}

func NewClient(conf *payloads.JwtConfig) *Client {
	return &Client{
		conf: conf,
	}
}

func (c *Client) Generate(userID int64) (string, error) {
	return c.GenerateWithExpiry(userID, c.conf.Expire)
}

func (c *Client) GenerateWithExpiry(userID int64, expiry time.Duration) (string, error) {
	if userID <= 0 {
		return "", errors.New("invalid user ID")
	}

	now := time.Now()
	exp := now.Add(expiry)

	claims := CustomClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    c.conf.Issuer,
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(exp),
			NotBefore: jwt.NewNumericDate(now),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString([]byte(c.conf.Secret))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return signedString, nil
}

func (c *Client) GenerateTokenPair(userID int64) (*TokenPair, error) {
	accessToken, err := c.Generate(userID)
	if err != nil {
		return nil, err
	}

	// Generate refresh token with longer expiry (7 days)
	refreshToken, err := c.GenerateWithExpiry(userID, 7*24*time.Hour)
	if err != nil {
		return nil, err
	}

	return &TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    time.Now().Add(c.conf.Expire),
	}, nil
}

func (c *Client) Parse(tokenString string) (*CustomClaims, error) {
	if tokenString == "" {
		return nil, ErrTokenNotFound
	}

	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(c.conf.Secret), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrExpiredToken
		}
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return nil, ErrInvalidSignature
		}
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrInvalidToken
}

// Validate validate
func (c *Client) Validate(tokenString string) error {
	_, err := c.Parse(tokenString)
	return err
}

func (c *Client) GetUserID(tokenString string) (int64, error) {
	claims, err := c.Parse(tokenString)
	if err != nil {
		return 0, err
	}
	return claims.UserID, nil
}

func (c *Client) IsExpired(tokenString string) bool {
	claims, err := c.Parse(tokenString)
	if err != nil {
		return true
	}
	return claims.ExpiresAt.Time.Before(time.Now())
}

func (c *Client) RefreshToken(refreshTokenString string) (*TokenPair, error) {
	claims, err := c.Parse(refreshTokenString)
	if err != nil {
		return nil, fmt.Errorf("invalid refresh token: %w", err)
	}

	// Generate new token pair
	return c.GenerateTokenPair(claims.UserID)
}

func (c *Client) GetRemainingTime(tokenString string) (time.Duration, error) {
	claims, err := c.Parse(tokenString)
	if err != nil {
		return 0, err
	}

	remaining := claims.ExpiresAt.Time.Sub(time.Now())
	if remaining < 0 {
		return 0, ErrExpiredToken
	}

	return remaining, nil
}

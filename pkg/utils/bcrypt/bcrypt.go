package bcrypt

import (
	"golang.org/x/crypto/bcrypt"
)

var config = Config{
	cost: 10,
}

type Opt func(*Config)
type Config struct {
	cost int
}

func WithCost(cost int) Opt {
	return func(c *Config) {
		c.cost = cost
	}
}

func Init(opts ...Opt) {
	for _, opt := range opts {
		opt(&config)
	}
}

func GeneratePassword(raw string) (string, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(raw), config.cost)
	if err != nil {
		return "", err
	}
	return string(password), nil
}
func ValidatePassword(hashed, origin string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(origin))
	if err != nil {
		return false
	}
	return true
}

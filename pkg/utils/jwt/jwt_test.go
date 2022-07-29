package jwt

import (
	"testing"

	"github.com/go-playground/assert/v2"

	"github.com/Xwudao/neter-template/pkg/config"
)

func TestJwtClient(t *testing.T) {
	koanf, err := config.NewConfig()
	assert.Equal(t, nil, err)
	client := NewClient(koanf)
	token, err := client.Generate(1)
	assert.Equal(t, nil, err)
	t.Logf("token: %s\n", token)
	assert.Equal(t, nil, client.Validate(token))

}

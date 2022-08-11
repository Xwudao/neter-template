package jwt

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Xwudao/neter-template/internal/core"
)

func TestJwtClient(t *testing.T) {
	app, cleanup, err := core.TestApp()
	defer cleanup()
	assert.Equal(t, nil, err)

	client := NewClient(app.Conf)
	token, err := client.Generate(1)
	assert.Equal(t, nil, err)
	t.Logf("token: %s\n", token)
	parse, err := client.Parse(token)
	assert.Equal(t, nil, err)
	t.Logf("parse: %+v\n", parse)
	assert.Equal(t, nil, client.Validate(token))

}

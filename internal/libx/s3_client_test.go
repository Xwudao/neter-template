package libx

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Xwudao/neter-template/internal/domain/payloads"
)

func TestNewS3Client(t *testing.T) {
	client, err := NewS3Client(&payloads.ProxyConfig{
		Username: "",
		Password: "",
		Addr:     "http://127.0.0.1:10809",
	}, &payloads.S3Config{})

	assert.Nil(t, err)
	assert.NotNil(t, client)
}

func TestS3Client_UploadBytes(t *testing.T) {
	client, err := NewS3Client(&payloads.ProxyConfig{
		Username: "",
		Password: "",
		Addr:     "http://127.0.0.1:10809",
	}, &payloads.S3Config{})
	assert.Nil(t, err)
	assert.NotNil(t, client)

	err = client.UploadBytes("v2fd-icons", "test.txt", []byte("test"))
	assert.Nil(t, err)
}

func TestS3Client_DownloadTo(t *testing.T) {
	client, err := NewS3Client(&payloads.ProxyConfig{
		Username: "",
		Password: "",
		Addr:     "http://127.0.0.1:10809",
	}, &payloads.S3Config{})
	assert.Nil(t, err)
	assert.NotNil(t, client)

	var buf bytes.Buffer

	err = client.DownloadTo("v2fd-icons", "test.txt", &buf)
	assert.Nil(t, err)

	assert.Equal(t, "test", buf.String())
}

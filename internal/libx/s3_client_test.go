package libx

import (
	"bytes"
	"os"
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
	}, s3TestConfig(t))
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
	}, s3TestConfig(t))
	assert.Nil(t, err)
	assert.NotNil(t, client)

	var buf bytes.Buffer

	err = client.DownloadTo("v2fd-icons", "test.txt", &buf)
	assert.Nil(t, err)

	assert.Equal(t, "test", buf.String())
}

// s3TestConfig returns credentials from the environment and skips the test when
// they are absent, so a fresh clone has a green test suite.
func s3TestConfig(t *testing.T) *payloads.S3Config {
	t.Helper()
	cfg := &payloads.S3Config{
		AccessKey: os.Getenv("S3_ACCESS_KEY"),
		SecretKey: os.Getenv("S3_SECRET_KEY"),
		Endpoint:  os.Getenv("S3_ENDPOINT"),
		Bucket:    os.Getenv("S3_BUCKET"),
	}
	if cfg.AccessKey == "" || cfg.SecretKey == "" {
		t.Skip("set S3_ACCESS_KEY/S3_SECRET_KEY/S3_ENDPOINT to run the S3 integration test")
	}
	return cfg
}

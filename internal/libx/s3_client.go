package libx

import (
	"bytes"
	"io"
	"net/http"
	"net/url"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/Xwudao/neter-template/internal/domain/payloads"
	"github.com/Xwudao/neter-template/pkg/utils"
)

type S3Client struct {
	svc *s3.S3
}

func NewS3Client(proxy *payloads.ProxyConfig, s3Config *payloads.S3Config) (*S3Client, error) {
	httpClient := http.DefaultClient
	var (
		accessKey = s3Config.AccessKey
		secretKey = s3Config.SecretKey
		endpoint  = s3Config.Endpoint
	)

	if proxy != nil {
		proxyStr := utils.BuildProxyUrl(proxy)
		proxyUrl, _ := url.Parse(proxyStr)
		transport := http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		}
		httpClient.Transport = &transport
	}

	sess, err := session.NewSession(&aws.Config{
		Credentials:      credentials.NewStaticCredentials(accessKey, secretKey, ""),
		Endpoint:         aws.String(endpoint),
		Region:           aws.String("auto"),
		DisableSSL:       aws.Bool(false),
		S3ForcePathStyle: aws.Bool(false),
		HTTPClient:       httpClient,
	})
	if err != nil {
		return nil, err
	}
	svc := s3.New(sess)

	return &S3Client{
		svc: svc,
	}, nil
}

// UploadFile 上传文件
func (s *S3Client) UploadFile(bucket, key string, buf io.ReadSeeker) error {
	_, err := s.svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   buf,
	})
	if err != nil {
		return err
	}

	return nil
}

// UploadBytes 上传字节
func (s *S3Client) UploadBytes(bucket, key string, data []byte) error {
	return s.UploadFile(bucket, key, bytes.NewReader(data))
}

// DownloadTo 下载到
func (s *S3Client) DownloadTo(bucket, key string, w io.Writer) error {
	resp, err := s.svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(w, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

// DownloadFile 下载文件
func (s *S3Client) DownloadFile(bucket, key string) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := s.DownloadTo(bucket, key, buf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

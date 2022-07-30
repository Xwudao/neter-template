package mail

import (
	"context"
	"crypto/tls"
	"time"

	"github.com/knadh/koanf"
	mail "github.com/xhit/go-simple-mail/v2"
	"go.uber.org/zap"

	"github.com/Xwudao/neter-template/internal/core"
)

type Config struct {
	//mail server
	host       string
	port       int
	username   string
	password   string
	encryption mail.Encryption

	//config
	authentication mail.AuthType
	keepAlive      bool
	connectTimeout time.Duration
	sendTimeout    time.Duration
	tlsConfig      *tls.Config
}

func NewConfig(host string, port int, username string, password string, keepAlive bool, connectTimeout time.Duration, sendTimeout time.Duration) *Config {
	c := &Config{host: host, port: port, username: username, password: password, keepAlive: keepAlive, connectTimeout: connectTimeout, sendTimeout: sendTimeout}

	c.authentication = mail.AuthPlain
	c.encryption = mail.EncryptionNone
	c.tlsConfig = &tls.Config{InsecureSkipVerify: true}

	return c
}

func NewConfigWithConf(conf *koanf.Koanf) *Config {
	host := conf.String("mail.host")
	port := conf.Int("mail.port")
	username := conf.String("mail.username")
	password := conf.String("mail.password")

	keepAlive := conf.Bool("mail.keepAlive")
	connectTimeout := conf.Duration("mail.connectTimeout")
	sendTimeout := conf.Duration("mail.sendTimeout")

	c := NewConfig(host, port, username, password, keepAlive, connectTimeout, sendTimeout)
	return c
}

type Message struct {
	From    string
	To      []string
	CC      []string
	Subject string
	//text body, html body can only set one
	TextBody string
	//text body, html body can only set one
	HtmlBody string
}

type Client struct {
	//client
	client  *mail.SMTPClient
	ctx     context.Context
	cancel  context.CancelFunc
	msgChan chan *Message

	logger *zap.SugaredLogger
}

func NewClientWithConf(ctx *core.AppContext, logger *zap.SugaredLogger, config *Config) (*Client, error) {
	server := mail.NewSMTPClient()

	// SMTP Server
	server.Host = config.host
	server.Port = config.port
	server.Username = config.username
	server.Password = config.password
	server.Encryption = config.encryption

	// Since v2.3.0 you can specified authentication type:
	// - PLAIN (default)
	// - LOGIN
	// - CRAM-MD5
	// server.Authentication = mail.AuthPlain

	// Variable to keep alive connection
	server.KeepAlive = config.keepAlive

	// Timeout for connect to SMTP Server
	server.ConnectTimeout = config.connectTimeout

	// Timeout for send the data and wait respond
	server.SendTimeout = config.sendTimeout

	// Set TLSConfig to provide custom TLS configuration. For example,
	// to skip TLS verification (useful for testing):
	server.TLSConfig = config.tlsConfig

	// SMTP client
	smtpClient, err := server.Connect()
	if err != nil {
		return nil, err
	}

	c := &Client{
		ctx:     ctx.Ctx,
		cancel:  ctx.Cancel,
		msgChan: make(chan *Message, 100),
		client:  smtpClient,
		logger:  logger,
	}

	return c, nil
}

func (c *Client) task() {
	log := c.logger
	for {
		select {
		case <-c.ctx.Done():
			log.Info("mail service stopped")
			return
		case m := <-c.msgChan:
			email := mail.NewMSG()
			//From Example <nube@example.com>
			email.SetFrom(m.From).
				AddTo(m.To...).
				AddCc(m.CC...).
				SetSubject(m.Subject)
			if m.HtmlBody != "" {
				email.SetBody(mail.TextHTML, m.HtmlBody)
			} else {
				email.SetBody(mail.TextPlain, m.TextBody)
			}

			// also you can add body from []byte with SetBodyData, example:
			// email.SetBodyData(mail.TextHTML, []byte(htmlBody))

			// add inline
			//email.Attach(&mail.File{FilePath: "/path/to/image.png", Name: "Gopher.png", Inline: true})

			// always check error after send
			if email.Error != nil {
				log.Error(email.Error)
				return
			}
			err := email.Send(c.client)
			if err != nil {
				log.Errorf("send email err: %s", err.Error())
				return
			}
		}
	}
}
func (c *Client) SendEmail(m *Message) {
	c.msgChan <- m
}

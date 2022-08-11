package mail

import (
	"testing"
)

func TestClient_SendEmail(t *testing.T) {
	//app, cleanup, err := core.TestApp()
	//if err != nil {
	//	t.Fatal(err)
	//}
	//defer cleanup()
	//config := NewConfigWithConf(app.Conf)
	//mailClient, err := NewClientWithConf(app.Ctx, app.Logger, config)
	//
	//assert.Nil(t, err)
	//
	//from := app.Conf.String("mail.from")
	//to := app.Conf.String("mail.to")
	//
	//assert.NotEmpty(t, from, "mail.from is empty")
	//assert.NotEmpty(t, to, "mail.to is empty")
	//
	//msg := &Message{
	//	From:     from,
	//	To:       []string{to},
	//	CC:       nil,
	//	Subject:  "hello wudao",
	//	TextBody: "I am your friend.",
	//	HtmlBody: "",
	//}
	//
	//mailClient.SendEmail(msg)
	//
	//time.Sleep(time.Second * 20)
}

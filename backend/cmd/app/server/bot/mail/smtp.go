package mail

import (
	"errors"
	"fmt"
	"github.com/Akvicor/glog"
	"github.com/labstack/echo/v4"
	"github.com/wneessen/go-mail"
	"msg/cmd/app/server/bot/mail/dto"
	"msg/cmd/app/server/bot/status"
	"msg/cmd/app/server/common/resp"
	"time"
)

type SMTPModel struct {
	SmtpHost string `toml:"smtp-host"`
	SmtpPort int    `toml:"smtp-port"`
	From     string `toml:"from"`
	Username string `toml:"username"`
	Password string `toml:"password"`
}

func NewSMTPModel(smtpHost string, smtpPort int, from, username, password string) *SMTPModel {
	return &SMTPModel{
		SmtpHost: smtpHost,
		SmtpPort: smtpPort,
		From:     from,
		Username: username,
		Password: password,
	}
}

const (
	TypeTextPlain = mail.TypeTextPlain
	TypeTextHtml  = mail.TypeTextHTML
)

func (m *SMTPModel) Status() status.SenderStatus {
	return status.SenderStatusOK
}

func (m *SMTPModel) Send(to, subject string, contentType mail.ContentType, body string) error {
	var err error

	msg := mail.NewMsg()
	err = msg.From(m.From)
	if err != nil {
		glog.Warning("邮件发送失败: %v", err)
		return errors.New("邮件发送失败")
	}
	err = msg.To(to)
	if err != nil {
		glog.Warning("邮件发送失败: %v", err)
		return errors.New("邮件发送失败")
	}
	msg.Subject(subject)
	msg.SetBodyString(contentType, body)

	try := 0
	for try < 3 {
		var client *mail.Client
		client, err = mail.NewClient(m.SmtpHost, mail.WithPort(m.SmtpPort), mail.WithSMTPAuth(mail.SMTPAuthPlain), mail.WithUsername(m.Username), mail.WithPassword(m.Password), mail.WithSSL())
		if err != nil {
			err = fmt.Errorf("邮件发送失败: %v", err)
			try++
			time.Sleep(time.Second)
			continue
		}

		err = client.DialAndSend(msg)
		_ = client.Close()
		if err != nil {
			err = fmt.Errorf("邮件发送失败: %v", err)
			try++
			time.Sleep(time.Second)
			continue
		}
		return nil
	}

	return err
}

// APISend 发送消息
func (m *SMTPModel) APISend(c echo.Context) (err error) {
	input := new(dto.SendModel)
	if err = c.Bind(input); err != nil {
		return resp.FailWithMsg(c, resp.BadRequest, fmt.Sprintf("错误输入: %v", err))
	}
	err = input.Validate()
	if err != nil {
		return resp.FailWithMsg(c, resp.BadRequest, fmt.Sprintf("错误输入: %v", err))
	}
	err = m.Send(input.To, input.Subject, input.Type, input.Msg)
	if err != nil {
		return resp.FailWithMsg(c, resp.BadRequest, fmt.Sprintf("%v", err))
	}
	return resp.Success(c)
}

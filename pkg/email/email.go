package email

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

type MailerOptions struct {
	SenderAlias string
}

func (mo *MailerOptions) WithSenderName(senderName string) *MailerOptions {
	mo.SenderAlias = senderName
	return mo
}

type Mailer struct {
	inner  *gomail.Dialer
	option *MailerOptions
}

func NewMailer(host string, port int, user, password string) *Mailer {
	dialer := gomail.NewDialer(host, port, user, password) // 发送邮件服务器、端口、发件人账号、发件人密码
	return &Mailer{
		inner: dialer,
		option: &MailerOptions{
			SenderAlias: user,
		},
	}
}

func (m *Mailer) SetOption(o *MailerOptions) *Mailer {
	m.option = o
	return m
}

func (m *Mailer) Send(receivers []string, subject, body string, attachFiles []string) error {
	msg := gomail.NewMessage()
	from := fmt.Sprintf("%v <%v>", m.option.SenderAlias, m.inner.Username)
	msg.SetHeader("From", from)
	msg.SetHeader("To", receivers...)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body)

	for _, v := range attachFiles {
		msg.Attach(v)
	}

	return m.inner.DialAndSend(msg)
}

package mail

import (
	"crypto/tls"
	"github.com/jordan-wright/email"
	"net/smtp"
)

type SendMailParams struct {
	From       string
	To         []string
	BCC        []string
	CC         []string
	Subject    string
	Addr       string
	Username   string
	Password   string
	Host       string
	ServerName string
}

func SendEmail(p SendMailParams, text string) error {
	e := email.NewEmail()
	e.From = p.From
	e.To = p.To
	e.Bcc = p.BCC
	e.Cc = p.CC
	e.Subject = p.Subject
	e.Text = []byte(text)
	err := e.SendWithTLS(p.Addr, smtp.PlainAuth("", p.Username, p.Password, p.Host),
		&tls.Config{
			InsecureSkipVerify: true,
			ServerName:         p.ServerName,
		})
	if err != nil {
		return err
	}
	return nil
}

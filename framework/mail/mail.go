package mail

import (
	"fmt"
	"net/smtp"
	"strings"

	"go-framework/framework/config"
)

type Mailer struct {
	config *config.Config
}

func NewMailer(cfg *config.Config) *Mailer {
	return &Mailer{config: cfg}
}

func (m *Mailer) Send(to, subject, body string) error {
	host := m.config.Get("MAIL_HOST", "localhost")
	port := m.config.Get("MAIL_PORT", "587")
	username := m.config.Get("MAIL_USERNAME", "")
	password := m.config.Get("MAIL_PASSWORD", "")
	from := m.config.Get("MAIL_FROM_ADDRESS", "app@localhost")

	addr := fmt.Sprintf("%s:%s", host, port)
	msg := []byte(fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s", from, to, subject, body))

	var auth smtp.Auth
	if username != "" {
		auth = smtp.PlainAuth("", username, password, host)
	}

	return smtp.SendMail(addr, auth, from, strings.Split(to, ","), msg)
}

func (m *Mailer) SendHTML(to, subject, htmlBody string) error {
	host := m.config.Get("MAIL_HOST", "localhost")
	port := m.config.Get("MAIL_PORT", "587")
	username := m.config.Get("MAIL_USERNAME", "")
	password := m.config.Get("MAIL_PASSWORD", "")
	from := m.config.Get("MAIL_FROM_ADDRESS", "app@localhost")

	addr := fmt.Sprintf("%s:%s", host, port)
	header := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\nMIME-Version: 1.0\r\nContent-Type: text/html; charset=UTF-8\r\n\r\n", from, to, subject)
	msg := []byte(header + htmlBody)

	var auth smtp.Auth
	if username != "" {
		auth = smtp.PlainAuth("", username, password, host)
	}

	return smtp.SendMail(addr, auth, from, strings.Split(to, ","), msg)
}

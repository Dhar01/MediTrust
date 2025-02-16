package utils

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

type EmailSender struct {
	dialer *gomail.Dialer
	from   string
	domain string
}

func NewEmailSender(username, password, from, domain, host string, port int) *EmailSender {
	dialer := gomail.NewDialer(host, port, username, password)
	return &EmailSender{
		dialer: dialer,
		from:   from,
		domain: domain,
	}
}

type EmailOptions struct {
	To            string
	Verification  bool
	ResetPassword bool
	FirstName     string
	Token         string
}

func (e *EmailSender) SendEmail(opts EmailOptions) error {
	switch {
	case opts.Verification:
		return e.SendVerificationEmail(opts)
	default:
		return fmt.Errorf("invalid email type")
	}
}

func (e *EmailSender) SendVerificationEmail(opts EmailOptions) error {
	subject := "User Verification Email - MediTrust"
	body := fmt.Sprintf(`
		<h3>Hello %s,</h3>
		<p>To verify your email, click here: <a href="http://%s:%d/api/v1/verify?token=%s">Verify Email</a></p>
	`, opts.FirstName, e.domain, e.dialer.Port, opts.Token)

	return e.sendEmail(opts.To, subject, body)
}

func (e *EmailSender) sendEmail(to, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", e.from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	return e.dialer.DialAndSend(m)
}

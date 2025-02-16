package utils

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

type EmailSender struct {
	dialer *gomail.Dialer
	from   string
}

func NewEmailSender(username, password, from, host string, port int) *EmailSender {
	dialer := gomail.NewDialer(host, port, username, password)
	return &EmailSender{
		dialer: dialer,
		from:   from,
	}
}

type EmailOptions struct {
	To            string
	Verification  bool
	ResetPassword bool
	FirstName     string
	Domain        string
	DomainPort    string
	Token         string
}

func (e *EmailSender) SendEmail(opts EmailOptions) error {
	switch {
	case opts.Verification:
		return e.SendVerificationEmail(opts)
	case opts.ResetPassword:
		return e.SendResetPasswordEmail(opts)
	default:
		return fmt.Errorf("invalid email type")
	}
}

func (e *EmailSender) SendVerificationEmail(opts EmailOptions) error {
	subject := "User Verification Email - MediTrust"
	body := fmt.Sprintf(`
		<h3>Hello %s,</h3>
		<p>To verify your email, click here: <a href="http://%s:%s/api/v1/verify?token=%s">Verify Email</a></p>
	`, opts.FirstName, opts.Domain, opts.DomainPort, opts.Token)

	return e.sendEmail(opts.To, subject, body)
}

func (e *EmailSender) SendResetPasswordEmail(opts EmailOptions) error {
	subject := "Reset your password"
	body := fmt.Sprintf(`
<h3>Hello %s,</h3>
<p>To reset your password, click here: <a href="http://%s:%s/api/v1/users/reset?token=%s">Reset password</a></p>
	`, opts.FirstName, opts.Domain, opts.DomainPort, opts.Token)

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

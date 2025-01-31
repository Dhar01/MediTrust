package utils

import (
	"fmt"
	"medicine-app/models"

	"gopkg.in/gomail.v2"
)

func SendVerificationEmail(userEmail, firstName, domain, token string, port int) error {
	m := gomail.NewMessage()
	m.SetHeader("From", models.EmailAdmin)
	m.SetHeader("To", userEmail)
	m.SetHeader("Subject", "User Verification Email")

	m.SetBody("text/html", fmt.Sprintf(`
		<h3>Hello %s,</h3>
		<p>To verify your email, click here: <a href="http://%s:%d/api/v1/verify?token=%s">Verify Email</a></p>
	`, firstName, domain, port, token))

	d := gomail.NewDialer(models.SMTPServer, models.SMTPPort, models.EmailAdmin, models.EmailAdminPass)

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

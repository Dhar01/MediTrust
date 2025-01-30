package utils

import (
	"log"
	"medicine-app/models"

	"gopkg.in/gomail.v2"
)

func SendVerificationEmail(userEmail string, token string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", models.BackendEmail)
	m.SetHeader("To", userEmail)
	m.SetHeader("Subject", "Verification Email")

	d := gomail.NewDialer(models.Host, models.Port, models.Username, models.Password)

	if err := d.DialAndSend(m); err != nil {
		log.Printf("Can't send email: %v", err)
	}

	return nil
}

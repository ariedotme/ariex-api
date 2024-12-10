package services

import (
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

type EmailService struct {
	SMTPHost string
	SMTPPort string
	Username string
	Password string
}

func (e *EmailService) Send(to, subject, message string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("SMTP_USERNAME"))
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", message)

	port, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		return err
	}

	d := gomail.NewDialer(
		os.Getenv("SMTP_HOST"),
		port,
		os.Getenv("SMTP_USERNAME"),
		os.Getenv("SMTP_PASSWORD"),
	)

	return d.DialAndSend(m)
}

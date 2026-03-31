package environment

import (
	"os"
)

func GetEmailEnv() (string, string, string, string, string) {
	fromEmail := os.Getenv("EMAIL_FROM_EMAIL")
	fromName := os.Getenv("EMAIL_FROM_NAME")
	password := os.Getenv("EMAIL_PASSWORD")
	smtpHost := os.Getenv("EMAIL_SMTP_HOST")
	smtpPort := os.Getenv("EMAIL_SMTP_PORT")

	return fromEmail, fromName, password, smtpHost, smtpPort
}

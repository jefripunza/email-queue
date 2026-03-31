package environment

import (
	"os"
	"strconv"
)

func GetEmailEnv() (string, string, string, string, string) {
	fromEmail := os.Getenv("EMAIL_FROM_EMAIL")
	fromName := os.Getenv("EMAIL_FROM_NAME")
	password := os.Getenv("EMAIL_PASSWORD")
	smtpHost := os.Getenv("EMAIL_SMTP_HOST")
	smtpPort := os.Getenv("EMAIL_SMTP_PORT")

	return fromEmail, fromName, password, smtpHost, smtpPort
}

func GetEmailSkipTLSVerify() bool {
	v := os.Getenv("EMAIL_SKIP_TLS_VERIFY")
	if v == "" {
		return false
	}
	b, err := strconv.ParseBool(v)
	if err != nil {
		return false
	}
	return b
}

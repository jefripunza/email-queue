package function

import (
	"crypto/tls"
	"email-queue/environment"
	"net/smtp"
)

func SendEmail(to string, subject string, body string) error {
	fromEmail, fromName, password, smtpHost, smtpPort := environment.GetEmailEnv()

	addr := smtpHost + ":" + smtpPort
	msg := []byte(
		"From: " + fromName + " <" + fromEmail + ">\r\n" +
			"To: " + to + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"MIME-Version: 1.0\r\n" +
			"Content-Type: text/html; charset=\"UTF-8\"\r\n\r\n" +
			body,
	)

	auth := smtp.PlainAuth("", fromEmail, password, smtpHost)
	client, err := smtp.Dial(addr)
	if err != nil {
		return err
	}

	tlsConfig := &tls.Config{
		ServerName: smtpHost,
	}
	if environment.GetEmailSkipTLSVerify() {
		tlsConfig.InsecureSkipVerify = true
	}
	if err = client.StartTLS(tlsConfig); err != nil {
		return err
	}
	if err = client.Auth(auth); err != nil {
		return err
	}
	if err = client.Mail(fromEmail); err != nil {
		return err
	}
	if err = client.Rcpt(to); err != nil {
		return err
	}
	w, err := client.Data()
	if err != nil {
		return err
	}
	_, err = w.Write(msg)
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	client.Quit()

	return nil
}

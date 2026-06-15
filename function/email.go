package function

import (
	"crypto/tls"
	"email-queue/environment"
	"io"
	"log"
	"net/smtp"
	"strings"
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

	client, err := smtp.Dial(addr)
	if err != nil {
		log.Printf("❌ [SMTP] Dial %s → %v", addr, err)
		return err
	}
	defer client.Quit()
	log.Printf("✅ [SMTP] Connected → %s", addr)

	tlsConfig := &tls.Config{ServerName: smtpHost}
	if environment.GetEmailSkipTLSVerify() {
		tlsConfig.InsecureSkipVerify = true
	}
	if err = client.StartTLS(tlsConfig); err != nil {
		log.Printf("❌ [SMTP] STARTTLS → %v", err)
		return err
	}
	log.Printf("✅ [SMTP] STARTTLS → TLS OK")

	auth := smtp.PlainAuth("", fromEmail, password, smtpHost)
	if err = client.Auth(auth); err != nil {
		log.Printf("❌ [SMTP] AUTH → %v", err)
		return err
	}
	log.Printf("✅ [SMTP] AUTH → OK")

	if err = client.Mail(fromEmail); err != nil {
		log.Printf("❌ [SMTP] MAIL FROM <%s> → %v", fromEmail, err)
		return err
	}
	log.Printf("✅ [SMTP] MAIL FROM <%s> → OK", fromEmail)

	if err = client.Rcpt(to); err != nil {
		log.Printf("❌ [SMTP] RCPT TO <%s> → %v", to, err)
		return err
	}
	log.Printf("✅ [SMTP] RCPT TO <%s> → OK", to)

	w, err := client.Data()
	if err != nil {
		log.Printf("❌ [SMTP] DATA → %v", err)
		return err
	}
	if _, err = io.Copy(w, strings.NewReader(string(msg))); err != nil {
		log.Printf("❌ [SMTP] Write body → %v", err)
		return err
	}
	if err = w.Close(); err != nil {
		log.Printf("❌ [SMTP] End DATA → %v", err)
		return err
	}
	log.Printf("✅ [SMTP] DATA → message accepted by server")

	log.Printf("✅ [SMTP] Email delivered: %s → %s", fromEmail, to)
	return nil
}

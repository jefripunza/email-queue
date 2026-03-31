package worker

import (
	"email-queue/environment"
	"email-queue/function"
	"email-queue/modules/message"
	"email-queue/variable"
	"log"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func StartEmailWorker() {
	go func() {
		for {
			delay := environment.GetDelay()
			delay_on_second := time.Duration(delay) * time.Second

			var msg message.EmailMessage
			if err := variable.Db.
				Session(&gorm.Session{Logger: logger.Default.LogMode(logger.Silent)}). // silent mode to avoid noise
				Where("is_sended = ? AND is_error = ?", false, false).
				First(&msg).
				Error; err != nil {
				time.Sleep(delay_on_second)
				continue
			}

			err := function.SendEmail(msg.To, msg.Subject, msg.Body)
			if err != nil {
				variable.Db.Model(&msg).Updates(map[string]interface{}{
					"is_error":      true,
					"error_message": err.Error(),
				})
				log.Printf("❌ Failed to send email (key=%s): %v", msg.Key, err)
			} else {
				variable.Db.Model(&msg).Updates(map[string]interface{}{
					"is_sended": true,
					"sended_at": time.Now(),
				})
				log.Printf("✅ Email sent successfully (key=%s)", msg.Key)
			}

			time.Sleep(delay_on_second)
		}
	}()
	log.Println("📧 Email worker started")
}

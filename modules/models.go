package modules

import (
	"email-queue/modules/message"

	"gorm.io/gorm"
)

func Models() []interface{} {
	return []interface{}{
		&message.EmailMessage{},
	}
}

func SeedAll(db *gorm.DB) {
	// message.EmailMessage{}.Seed(db)
}

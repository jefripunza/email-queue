package message

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EmailMessage struct {
	ID           uuid.UUID `json:"id" gorm:"type:char(36);primaryKey"`
	Key          string    `json:"key" gorm:"type:varchar(50);not null;uniqueIndex"`
	To           string    `json:"to" gorm:"not null"` // email
	Subject      string    `json:"subject" gorm:"not null"`
	Body         string    `json:"body" gorm:"type:text;not null"` // text/html
	IsSended     bool      `json:"is_sended" gorm:"default:false"`
	IsError      bool      `json:"is_error" gorm:"default:false"`
	ErrorMessage string    `json:"error_message" gorm:""` // if error
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	SendedAt     time.Time `json:"sended_at" gorm:"autoCreateTime"`
}

func (s *EmailMessage) BeforeCreate(tx *gorm.DB) error {
	if s.ID == uuid.Nil {
		s.ID = uuid.New()
	}
	return nil
}

func (s *EmailMessage) Map() map[string]any {
	return map[string]any{
		"id":            s.ID,
		"key":           s.Key,
		"to":            s.To,
		"subject":       s.Subject,
		"body":          s.Body,
		"is_sended":     s.IsSended,
		"is_error":      s.IsError,
		"error_message": s.ErrorMessage,
		"created_at":    s.CreatedAt,
		"sended_at":     s.SendedAt,
	}
}

func (EmailMessage) Seed(db *gorm.DB) {
	// TODO: Seed default email messages
}

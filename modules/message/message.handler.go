package message

import (
	"email-queue/dto"
	"email-queue/function"
	"email-queue/variable"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type SendEmailRequest struct {
	To      string `json:"to" validate:"required"`
	Subject string `json:"subject" validate:"required"`
	Body    string `json:"body" validate:"required"`
}

func SendEmail(c *fiber.Ctx) error {
	var body SendEmailRequest
	if err := function.RequestBody(c, &body); err != nil {
		return dto.BadRequest(c, err.Error(), nil)
	}

	key := uuid.New().String()
	msg := EmailMessage{
		Key:     key,
		To:      body.To,
		Subject: body.Subject,
		Body:    body.Body,
	}
	if err := variable.Db.
		Create(&msg).
		Error; err != nil {
		return dto.InternalServerError(c, "Failed to queue email", nil)
	}

	return dto.Created(c, "Email queued successfully", fiber.Map{
		"key": key,
	})
}

func CheckStatus(c *fiber.Ctx) error {
	key := c.Params("key")
	if key == "" {
		return dto.BadRequest(c, "Key is required", nil)
	}

	var msg EmailMessage
	if err := variable.Db.Where("key = ?", key).First(&msg).Error; err != nil {
		return dto.NotFound(c, "Email message not found", nil)
	}

	return dto.OK(c, "Email status retrieved", msg.Map())
}

package message

import (
	"github.com/gofiber/fiber/v2"
)

func Route(api fiber.Router) {
	api.Post("/send", SendEmail)
	api.Get("/status/:key", CheckStatus)
}

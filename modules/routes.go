package modules

import (
	"email-queue/middlewares"
	"email-queue/modules/message"

	"github.com/gofiber/fiber/v2"
)

func Routes(app fiber.Router) {
	api := app.Group("/api")

	// Message
	message.Route(api.Group("/message", middlewares.UseApiKey))

}

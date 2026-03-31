package middlewares

import (
	"email-queue/environment"

	"github.com/gofiber/fiber/v2"
)

func UseApiKey(c *fiber.Ctx) error {
	apiKey := c.Get("X-API-Key")
	if apiKey != environment.GetApiKey() {
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}
	return c.Next()
}

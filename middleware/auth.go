package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddle(c *fiber.Ctx) error {
	token := c.Cookies("token")

	if token != os.Getenv("PASSWORD") {
		return c.Status(401).JSON(fiber.Map{"error": "unauthorized"})
	}

	return c.Next()
}

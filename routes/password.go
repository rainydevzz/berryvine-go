package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Password(app *fiber.App) {
	app.Get("/password", func(c *fiber.Ctx) error {
		return c.Render("password", fiber.Map{})
	})
}

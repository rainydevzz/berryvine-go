package routes

import (
	"berryvine/middleware"

	"github.com/gofiber/fiber/v2"
)

func Upload(app *fiber.App) {
	app.Get("/upload", middleware.AuthMiddle, func(c *fiber.Ctx) error {
		return c.Render("upload", fiber.Map{})
	})
}

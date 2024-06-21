package routes

import (
	"berryvine/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Root(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		var metadata []models.Metadata
		result, _ := c.Locals("db").(*gorm.DB).Find(&metadata).Rows()

		defer result.Close()
		return c.Render("index", fiber.Map{"files": metadata})
	})
}

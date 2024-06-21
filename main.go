package main

import (
	"berryvine/models"

	"berryvine/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/gofiber/template/handlebars/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("main.db"), &gorm.Config{})

	db.AutoMigrate(&models.Metadata{})

	hbs := handlebars.New("./static", ".hbs")

	if err != nil {
		panic("failed to connect")
	}

	app := fiber.New(fiber.Config{
		Views: hbs,
	})

	app.Use(func(c *fiber.Ctx) error {
		c.Locals("db", db)
		return c.Next()
	})

	app.Static("/files", "./files")

	app.Get("/swagger/*", swagger.HandlerDefault)

	routes.Root(app)
	routes.Upload(app)
	routes.Password(app)
	routes.SubmitPassword(app)
	routes.SubmitUpload(app)

	app.Listen(":3000")
}

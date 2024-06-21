package routes

import (
	"berryvine/middleware"
	"berryvine/models"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SubmitUpload(app *fiber.App) {
	app.Post("/upload", middleware.AuthMiddle, func(c *fiber.Ctx) error {
		if form, err := c.MultipartForm(); err == nil {
			f := form.File["file"]
			for _, file := range f {
				if file.Header["Content-Type"][0] != "video/mp4" {
					return c.SendStatus(400)
				}
				c.SaveFile(file, fmt.Sprintf("./files/%s", file.Filename))
				c.Locals("db").(*gorm.DB).Create(models.Metadata{File: file.Filename, Name: "My Clip"})
			}
		}
		return c.Redirect("/", 200)
	})
}

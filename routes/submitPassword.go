package routes

import (
	"berryvine/models"
	"os"

	"github.com/gofiber/fiber/v2"
)

func SubmitPassword(app *fiber.App) {
	app.Post("/password", func(c *fiber.Ctx) error {
		PASSWORD := os.Getenv("PASSWORD")
		authForm := new(models.AuthForm)
		cookie := new(fiber.Cookie)

		if err := c.BodyParser(authForm); err != nil {
			return c.SendStatus(400)
		}

		if authForm.Auth == PASSWORD {
			cookie.Name = "token"
			cookie.Value = PASSWORD
			c.Cookie(cookie)
			return c.Redirect("/", 200)
		} else {
			return c.SendStatus(401)
		}
	})
}

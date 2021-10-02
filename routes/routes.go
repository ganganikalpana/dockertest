package routes

import (
	"test/controller"

	"github.com/gofiber/fiber/v2"
)

func SetUp(app *fiber.App) {
	//app.Get("/", controller.RedirectTo)
	app.Post("/v1/auth/users", controller.NewUser)

}

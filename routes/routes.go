package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/harshthakkar09/blog-website/controller"
)

func Setup(app *fiber.App) {
	app.Post("api/register", controller.Register)
}

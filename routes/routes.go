package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/harshthakkar09/blog-website/controller"
	"github.com/harshthakkar09/blog-website/middleware"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)

	app.Use(middleware.IsAuthanticate)
	app.Post("/api/post", controller.CreatePost)
	app.Get("/api/allpost", controller.AllPosts)
	app.Get("/api/allpost/:id", controller.DetailPost)
	app.Put("/api/updatepost/:id", controller.UpdatePost)
	app.Delete("/api/deletepost/:id", controller.DeletePost)
}

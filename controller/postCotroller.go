package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/harshthakkar09/blog-website/database"
	"github.com/harshthakkar09/blog-website/models"
)

func CreatePost(c *fiber.Ctx) error {
	var blogpost models.Blog
	if err := c.BodyParser(&blogpost); err != nil {
		fmt.Println("Unable to parase body ")
	}
	if err := database.DB.Create(&blogpost).Error; err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Invalid payload",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Congratulations!, Your post is live",
	})
}

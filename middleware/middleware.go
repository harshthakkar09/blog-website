package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/harshthakkar09/blog-website/util"
)

func IsAuthanticate(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	if _, err := util.ParseJwt(cookie); err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	return c.Next()
}

package routes

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterAdminRoutes(router fiber.Router) {

	admin := router.Group("/admin")

	admin.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Admin Home")
	})
}

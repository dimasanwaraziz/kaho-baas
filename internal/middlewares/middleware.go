package middlewares

import (
	"Kaho_BaaS/internal/response"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

// NewAuthMiddleware creates a middleware with the provided session store
func NewAuthMiddleware(store *session.Store) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(response.Error(err.Error(), nil))
		}

		// Periksa apakah pengguna sudah login
		if sess.Get("isLoggedIn") != true {
			return c.Status(fiber.StatusUnauthorized).JSON(response.Error("Unauthorized", nil))
		}

		// Lanjutkan ke handler berikutnya
		return c.Next()
	}
}

package handlers

import (
	"Kaho_BaaS/internal/response"

	"github.com/gofiber/fiber/v2"
)

// Logout
//
//	@Summary	Logout
//	@Tags		Account
//	@Accept		json
//	@Produce	json
//	@Router		/account/sessions/current [delete]
func (h *authHandler) Logout(c *fiber.Ctx) error {
	sess, err := h.session.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	sess.Destroy()

	return c.JSON(response.Success("Logout Successful", nil))
}

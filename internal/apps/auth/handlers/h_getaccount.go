package handlers

import (
	"Kaho_BaaS/internal/response"

	"github.com/gofiber/fiber/v2"
)

// GetAccount
//
//	@Summary	Get Detail Account
//	@Tags		Account
//	@Produce	json
//	@Router		/account [get]
func (h *authHandler) GetAccount(c *fiber.Ctx) error {
	sess, err := h.session.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	username := sess.Get("username")
	email := sess.Get("email")
	id := sess.Get("id")
	company_id := sess.Get("company_id")
	company_name := sess.Get("company_name")
	roles := sess.Get("roles")
	name := sess.Get("full_name")

	return c.JSON(response.Success("Get Account Successful", fiber.Map{
		"username": username,
		"email":    email,
		"id":       id,
		"company": fiber.Map{
			"id":   company_id,
			"name": company_name,
		},
		"roles":     roles,
		"full_name": name,
	}))
}

package handlers

import (
	"Kaho_BaaS/internal/response"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// Login
//
//	@Summary		Login
//	@Description	Untuk email, bisa menggunakan email atau username
//	@Tags			Account
//	@Accept			multipart/form-data
//	@Produce		json
//	@Param			email		formData	string	true	"Email or username"
//	@Param			password	formData	string	true	"Password"
//	@Router			/account/sessions/email [post]
func (h *authHandler) Login(c *fiber.Ctx) error {
	// Get form values
	email := c.FormValue("email")
	password := c.FormValue("password")

	if email == "" || password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(response.Error("Email and password are required", nil))
	}

	user, err := h.service.Login(email, password)
	if err != nil {
		fmt.Println(err)
		return c.Status(fiber.StatusUnauthorized).JSON(response.Error("Invalid email or password", nil))
	}

	// Ambil sesi dari store
	sess, err := h.session.Get(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	// Simpan data sesi
	sess.Set("id", user.ID)
	sess.Set("username", user.Username)
	sess.Set("email", user.Email)
	sess.Set("company_id", user.CompanyId)
	sess.Set("company_name", user.CompanyName)
	sess.Set("roles", user.Roles)
	sess.Set("full_name", user.Name)
	sess.Set("isLoggedIn", true)

	// Simpan sesi ke penyimpanan
	if err := sess.Save(); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	data := fiber.Map{
		"username":   user.Username,
		"email":      user.Email,
		"isLoggedIn": true,
	}

	return c.JSON(response.Success("Login Successful", data))
}

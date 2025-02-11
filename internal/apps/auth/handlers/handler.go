package handlers

import (
	"Kaho_BaaS/internal/apps/auth/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type AuthHandler interface {
	Login(c *fiber.Ctx) error
	Logout(c *fiber.Ctx) error
	GetAccount(c *fiber.Ctx) error
}

type authHandler struct {
	service services.AuthService
	session *session.Store
}

func NewAuthHandler(services services.AuthService, session *session.Store) AuthHandler {
	return &authHandler{
		service: services,
		session: session,
	}
}

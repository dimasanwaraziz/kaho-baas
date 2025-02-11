package routes

import (
	"Kaho_BaaS/internal/apps/auth/handlers"
	"Kaho_BaaS/internal/apps/auth/services"
	"Kaho_BaaS/internal/middlewares"
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

func RegisterAuthRoutes(router fiber.Router, db *sql.DB, store *session.Store) {

	services := services.NewAuthService(db, store)
	authHandler := handlers.NewAuthHandler(services, store)
	authMiddleware := middlewares.NewAuthMiddleware(store)

	auth := router.Group("/account")

	auth.Get("/", authMiddleware, authHandler.GetAccount)
	auth.Post("/sessions/email", authHandler.Login)
	auth.Delete("/sessions/current", authHandler.Logout)
}

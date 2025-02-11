package server

import (
	_ "Kaho_BaaS/docs"
	adminRoutes "Kaho_BaaS/internal/apps/admin/routes"
	authRoutes "Kaho_BaaS/internal/apps/auth/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

func (s *FiberServer) RegisterFiberRoutes() {
	// Apply CORS middleware
	s.App.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS,PATCH",
		AllowHeaders:     "Accept,Authorization,Content-Type",
		AllowCredentials: false, // credentials require explicit origins
		MaxAge:           300,
	}))

	s.App.Get("/docs/*", swagger.HandlerDefault)

	s.App.Use(logger.New())

	v1 := s.App.Group("/v1")

	adminRoutes.RegisterAdminRoutes(v1)
	authRoutes.RegisterAuthRoutes(v1, s.db.GetDB(), s.store)

	s.App.Get("/", s.HelloWorldHandler)

	s.App.Get("/health", s.healthHandler)

}

func (s *FiberServer) HelloWorldHandler(c *fiber.Ctx) error {
	resp := fiber.Map{
		"message": "Hello World",
	}

	return c.JSON(resp)
}

func (s *FiberServer) healthHandler(c *fiber.Ctx) error {
	return c.JSON(s.db.Health())
}

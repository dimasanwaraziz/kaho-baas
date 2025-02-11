package server

import (
	"log"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/postgres/v3"

	"Kaho_BaaS/internal/database"
)

type FiberServer struct {
	*fiber.App

	db    database.Service
	store *session.Store
}

func New() *FiberServer {

	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalf("Invalid port value: %v", err)
	}

	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "Kaho_BaaS",
			AppName:      "Kaho_BaaS",
		}),

		db: database.New(),
		store: session.New(session.Config{
			Expiration: 31536000, // 1 year in seconds (365 * 24 * 60 * 60)
			Storage: postgres.New(postgres.Config{
				Host:     os.Getenv("DB_HOST"),
				Port:     port,
				Username: os.Getenv("DB_USERNAME"),
				Password: os.Getenv("DB_PASSWORD"),
				Database: os.Getenv("DB_DATABASE"),
				Table:    "fiber_storage",
			}),
			CookieSameSite: "None",
		}),
	}

	return server
}

package services

import (
	"database/sql"

	"github.com/gofiber/fiber/v2/middleware/session"
)

type AuthService interface {
	Login(email, password string) (UserData, error)
}

type authService struct {
	db      *sql.DB
	session *session.Store
}

func NewAuthService(db *sql.DB, session *session.Store) AuthService {
	return &authService{
		db:      db,
		session: session,
	}
}

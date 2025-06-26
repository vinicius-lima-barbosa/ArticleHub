package database

import (
	"database/sql"

	"articlehub-api/internal/database/repository"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
)

// Service represents a service that interacts with a database.
type Service interface {
	UserRepo() repository.UserRepository

	// Health returns a map of health status information.
	// The keys and values in the map are service-specific.
	Health() map[string]string

	// Close terminates the database connection.
	// It returns an error if the connection cannot be closed.
	Close() error
}

type service struct {
	db       *sql.DB
	userRepo repository.UserRepository
}

// func (s *service) migrate() error {
// 	query := `
// 	CREATE TABLE IF NOT EXISTS users (
// 		id SERIAL PRIMARY KEY,
// 		name TEXT NOT NULL,
// 		email TEXT UNIQUE NOT NULL,
// 		password TEXT NOT NULL,
// 		created_at TIMESTAMPTZ DEFAULT now(),
// 		updated_at TIMESTAMPTZ DEFAULT now()
// 	);`
// 	_, err := s.db.Exec(query)
// 	return err
// }

func New() Service {
	db := NewConnection()
	return &service{
		db:       db,
		userRepo: repository.NewUserRepository(db),
	}
}

func (s *service) UserRepo() repository.UserRepository {
	return s.userRepo
}

func (s *service) Health() map[string]string {
	return Health(s.db)
}

func (s *service) Close() error {
	return s.db.Close()
}

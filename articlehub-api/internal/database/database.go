package database

import (
	"database/sql"

	user_repository "articlehub-api/internal/repository/user-repository"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
)

// Service represents a service that interacts with a database.
type Service interface {
	UserRepo() user_repository.UserRepository

	// Health returns a map of health status information.
	// The keys and values in the map are service-specific.
	Health() map[string]string

	// Close terminates the database connection.
	// It returns an error if the connection cannot be closed.
	Close() error
}

type service struct {
	db       *sql.DB
	userRepo user_repository.UserRepository
}

func New() Service {
	db := NewConnection()
	return &service{
		db:       db,
		userRepo: user_repository.NewUserRepository(db),
	}
}

func (s *service) UserRepo() user_repository.UserRepository {
	return s.userRepo
}

func (s *service) Health() map[string]string {
	return Health(s.db)
}

func (s *service) Close() error {
	return s.db.Close()
}

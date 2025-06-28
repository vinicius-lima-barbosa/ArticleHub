package repository

import (
	"context"
	"database/sql"
	"fmt"

	"articlehub-api/internal/model"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User) error
	GetUsers(ctx context.Context) ([]model.User, error)
	GetUserById(ctx context.Context, id int) (*model.User, error)
	UpdateUser(ctx context.Context, id int, user *model.User) error
	DeleteUser(ctx context.Context, id int) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	if db == nil {
		fmt.Println("❌ Banco de dados não inicializado")
	} else {
		fmt.Println("✅ Banco de dados conectado com sucesso")
	}
	return &userRepository{db: db}
}

// CreateUser creates a new user in the database.
func (r *userRepository) CreateUser(ctx context.Context, user *model.User) error {
	query := `INSERT INTO users (name, email, password, created_at, updated_at) VALUES ($1, $2, $3, NOW(), NOW()) RETURNING id, created_at, updated_at`
	err := r.db.QueryRowContext(ctx, query, user.Name, user.Email, user.Password).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func (r *userRepository) GetUsers(ctx context.Context) ([]model.User, error) {
	query := `SELECT id, name, email, created_at, updated_at FROM users ORDER BY created_at DESC`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *userRepository) GetUserById(ctx context.Context, id int) (*model.User, error) {
	query := `SELECT id, name, email, created_at, updated_at FROM users WHERE id = $1`
	var user model.User
	err := r.db.QueryRowContext(ctx, query, id).
		Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, id int, user *model.User) error {
	query := `UPDATE users SET name = $1, email = $2, updated_at = NOW() WHERE id = $3 RETURNING updated_at`
	return r.db.QueryRowContext(ctx, query, user.Name, user.Email, id).
		Scan(&user.UpdatedAt)
}

func (r *userRepository) DeleteUser(ctx context.Context, id int) error {
	query := `DELETE FROM users WHERE id = $1`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return fmt.Errorf("user not found")
	}
	return nil
}

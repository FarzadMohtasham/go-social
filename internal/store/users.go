package store

import (
	"context"
	"database/sql"
)

type User struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"-"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

type UsersStore struct {
	db *sql.DB
}

func (s *UsersStore) Create(ctx context.Context, user *User) error {
	query := `
		INSERT INTO users (username, password, email) VALUES ($1, $2, $3) RETURNING id, created_at
	`

	err := s.db.QueryRowContext(ctx, query, user.Username, user.Password, user.Email).Scan(
		&user.Username, &user.Password, &user.Email,
	)

	if err != nil {
		return err
	}

	return nil
}

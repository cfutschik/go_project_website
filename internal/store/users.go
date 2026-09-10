package store

import (
	"context"
	"database/sql"
)

type User struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	CreatedAt string `json:"created_at"`
}

type UserStore struct {
	db *sql.DB
}

func (u *UserStore) Create(ctx context.Context, user *User) error {
	query := `
		INSERT INTO users (username, password, email) VALUES($1, $2, $3) RETURNING id,
		created_at
	`

	if err := u.db.QueryRowContext(
		ctx,
		query,
		user.Username,
		user.Password,
		user.Email,
	).Scan(
		&user.ID,
		&user.CreatedAt,
	); err != nil {
		return err
	}

	return nil
}

func (u *UserStore) GetByID(ctx context.Context, id int64) (*User, error) {
	query := `SELECT * FROM users WHERE id = $1`

	var user User

	if err := u.db.QueryRowContext(
		ctx,
		query,
		id).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	); err != nil {
		return nil, err
	}

	return &user, nil

}

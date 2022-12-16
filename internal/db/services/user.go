package services

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

type UserService struct {
	DBPool *pgxpool.Pool
}

func (s *UserService) Get(userId string) (*User, error) {
	var u User

	row := s.DBPool.QueryRow(
		context.Background(),
		"SELECT id, email FROM users WHERE id = $1",
		userId,
	)

	if err := row.Scan(&u.ID, &u.Email); err != nil {
		return nil, err
	}

	return &u, nil
}

func (s *UserService) Insert(email, password string) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return nil, err
	}

	var u User

	row := s.DBPool.QueryRow(
		context.Background(),
		`INSERT INTO users (email, hashed_password) VALUES ($1, $2)
			RETURNING id, email`,
		email,
		hashedPassword,
	)

	if err := row.Scan(&u.ID, &u.Email); err != nil {
		return nil, err
	}

	return &u, nil
}

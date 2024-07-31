package db_repo

import (
	"context"
	"database/sql"
	"war_ticket/internal/domain"
	"war_ticket/internal/interfaces"
)

type UserRepositoryImpl struct {
	DB *sql.DB
}

type UserRepository interface {
	interfaces.Saver[domain.User]
	FindByApiKey(apiKey string) (*domain.User, error)
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &UserRepositoryImpl{
		DB: db,
	}
}

// Save implements UserRepository.
func (u *UserRepositoryImpl) Save(ctx context.Context, value *domain.User) (*domain.User, error) {

	query := `INSERT INTO users(api_key, username) VALUES ($1, $2)`

	_, err := u.DB.ExecContext(ctx, query, value.ApiKey, value.Username)

	if err != nil {
		return nil, err
	}

	return value, nil
}

// FindByApiKey implements UserRepository.
func (u *UserRepositoryImpl) FindByApiKey(apiKey string) (*domain.User, error) {

	var user domain.User

	query := `SELECT username, api_key FROM users WHERE api_key = $1 LIMIT 1`

	row := u.DB.QueryRow(query, apiKey)

	err := row.Scan(&user.Username, &user.ApiKey)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

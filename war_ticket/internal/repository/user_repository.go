package repository

import (
	"context"
	"war_ticket/internal/domain"
	errr "war_ticket/internal/err"
	"war_ticket/internal/interfaces"
)

type UserRepositoryImpl struct {
	Users map[string]domain.User
}

type UserRepository interface {
	interfaces.Saver[domain.User]
	FindByApiKey(apiKey string) (*domain.User, error)
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{
		Users: make(map[string]domain.User),
	}
}

// Save implements UserRepository.
func (u *UserRepositoryImpl) Save(ctx context.Context, value *domain.User) (*domain.User, error) {
	_, ok := u.Users[value.ApiKey]

	if ok {
		return nil, errr.ErrDuplicateID
	}

	u.Users[value.ApiKey] = *value

	return value, nil
}

// FindByApiKey implements UserRepository.
func (u *UserRepositoryImpl) FindByApiKey(apiKey string) (*domain.User, error) {
	value, ok := u.Users[apiKey]

	if !ok {
		return nil, errr.ErrNotFound
	}

	return &value, nil
}

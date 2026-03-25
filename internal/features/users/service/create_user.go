package users_service

import (
	"context"
	"fmt"

	domain "github.com/musashimiyomoto/todo-app/internal/core/domain"
)

func (s *UsersService) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	if err := user.Validate(); err != nil {
		return domain.User{}, fmt.Errorf("Validate user domain: %w", err)
	}

	userDomain, err := s.usersRepository.CreateUser(ctx, user)
	if err != nil {
		return domain.User{}, fmt.Errorf("Create user in repository: %w", err)
	}

	return userDomain, nil
}

package users_service

import (
	"context"
	"fmt"

	"github.com/musashimiyomoto/todo-app/internal/core/domain"
)

func (s *UsersService) CreateUser(ctx context.Context, user domain.User) (domain.User, error) {
	if err := user.Validate(); err != nil {
		return domain.User{}, fmt.Errorf("Validate user domain: %w", err)
	}

	user, err := s.usersRepository.CreateUser(ctx, user)
	if err != nil {
		return domain.User{}, fmt.Errorf("Create user in repository: %w", err)
	}

	return user, nil
}

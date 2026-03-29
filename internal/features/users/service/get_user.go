package users_service

import (
	"context"
	"fmt"

	domain "github.com/musashimiyomoto/todo-app/internal/core/domain"
)

func (s *UsersService) GetUser(ctx context.Context, id int) (domain.User, error) {
	userDomain, err := s.usersRepository.GetUser(ctx, id)
	if err != nil {
		return domain.User{}, fmt.Errorf("Get user from repository: %w", err)
	}

	return userDomain, nil
}

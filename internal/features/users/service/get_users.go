package users_service

import (
	"context"
	"fmt"

	domain "github.com/musashimiyomoto/todo-app/internal/core/domain"
	core_errors "github.com/musashimiyomoto/todo-app/internal/core/errors"
)

func (s *UsersService) GetUsers(ctx context.Context, limit *int, offset *int) ([]domain.User, error) {
	if limit != nil && *limit < 0 {
		return nil, fmt.Errorf(
			"Invalid 'limit' query param: must be a non-negative integer: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if offset != nil && *offset < 0 {
		return nil, fmt.Errorf(
			"Invalid 'offset' query param: must be a non-negative integer: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if limit != nil && offset != nil && *limit <= *offset {
		return nil, fmt.Errorf(
			"Invalid query params: 'limit' must be greater than 'offset': %w",
			core_errors.ErrInvalidArgument,
		)
	}

	userDomains, err := s.usersRepository.GetUsers(ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("Get users from repository: %w", err)
	}

	return userDomains, nil
}

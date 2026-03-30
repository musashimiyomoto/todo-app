package users_service

import (
	"context"
	"fmt"

	domain "github.com/musashimiyomoto/todo-app/internal/core/domain"
)

func (s *UsersService) PatchUser(ctx context.Context, id int, patch domain.UserPatch) (domain.User, error) {
	userDomain, err := s.usersRepository.GetUser(ctx, id)
	if err != nil {
		return domain.User{}, fmt.Errorf("Get user from repository: %w", err)
	}

	if err := userDomain.ApplyPatch(patch); err != nil {
		return domain.User{}, fmt.Errorf("Apply patch to user: %w", err)
	}

	patchedUser, err := s.usersRepository.PatchUser(ctx, id, userDomain)
	if err != nil {
		return domain.User{}, fmt.Errorf("Patch user in repository: %w", err)
	}

	return patchedUser, nil
}

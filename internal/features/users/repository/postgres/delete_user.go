package users_postgres_repository

import (
	"context"
	"fmt"

	core_errors "github.com/musashimiyomoto/todo-app/internal/core/errors"
)

func (r *UsersRepository) DeleteUser(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		DELETE FROM users
		WHERE id = $1;
	`

	cmdTag, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("Failed to execute delete query: %w", err)
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("Delete user with ID %d: %w", id, core_errors.ErrNotFound)
	}

	return nil
}

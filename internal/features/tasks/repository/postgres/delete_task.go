package tasks_postgres_repository

import (
	"context"
	"fmt"

	core_errors "github.com/musashimiyomoto/todo-app/internal/core/errors"
)

func (r *TasksRepository) DeleteTask(ctx context.Context, id int) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	query := `
		DELETE FROM tasks
		WHERE id = $1;
	`

	cmdTag, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("Failed to execute delete query: %w", err)
	}

	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf("Delete task with ID %d: %w", id, core_errors.ErrNotFound)
	}

	return nil
}

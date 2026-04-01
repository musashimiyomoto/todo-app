package tasks_service

import (
	"context"
	"fmt"

	domain "github.com/musashimiyomoto/todo-app/internal/core/domain"
	core_errors "github.com/musashimiyomoto/todo-app/internal/core/errors"
)

func (s *TasksService) GetTasks(ctx context.Context, userID *int, limit *int, offset *int) ([]domain.Task, error) {
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

	taskDomains, err := s.tasksRepository.GetTasks(ctx, userID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("Get tasks from repository: %w", err)
	}

	return taskDomains, nil
}

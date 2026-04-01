package tasks_service

import (
	"context"
	"fmt"

	domain "github.com/musashimiyomoto/todo-app/internal/core/domain"
)

func (s *TasksService) PatchTask(ctx context.Context, id int, patch domain.TaskPatch) (domain.Task, error) {
	taskDomain, err := s.tasksRepository.GetTask(ctx, id)
	if err != nil {
		return domain.Task{}, fmt.Errorf("Get task from repository: %w", err)
	}

	if err := taskDomain.ApplyPatch(patch); err != nil {
		return domain.Task{}, fmt.Errorf("Apply patch to task: %w", err)
	}

	patchedTask, err := s.tasksRepository.PatchTask(ctx, id, taskDomain)
	if err != nil {
		return domain.Task{}, fmt.Errorf("Patch task in repository: %w", err)
	}

	return patchedTask, nil
}

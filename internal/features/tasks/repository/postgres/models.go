package tasks_postgres_repository

import (
	"time"

	domain "github.com/musashimiyomoto/todo-app/internal/core/domain"
)

type TaskModel struct {
	ID           int
	Version      int
	Title        string
	Description  *string
	Completed    bool
	CreatedAt    time.Time
	CompletedAt  *time.Time
	AuthorUserID int
}

func taskDomainsFromModels(taskModels []TaskModel) []domain.Task {
	tasksDomains := make([]domain.Task, len(taskModels))

	for i, taskModel := range taskModels {
		tasksDomains[i] = domain.NewTask(
			taskModel.ID,
			taskModel.Version,
			taskModel.Title,
			taskModel.Description,
			taskModel.Completed,
			taskModel.CreatedAt,
			taskModel.CompletedAt,
			taskModel.AuthorUserID,
		)
	}

	return tasksDomains
}

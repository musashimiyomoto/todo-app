package tasks_transport_http

import (
	"time"

	domain "github.com/musashimiyomoto/todo-app/internal/core/domain"
)

type TaskDTOResponse struct {
	ID           int        `json:"id"`
	Version      int        `json:"version"`
	Title        string     `json:"title"`
	Description  *string    `json:"description"`
	Completed    bool       `json:"completed"`
	CreatedAt    time.Time  `json:"created_at"`
	CompletedAt  *time.Time `json:"completed_at"`
	AuthorUserID int        `json:"author_user_id"`
}

func taskDTOFromDomain(domain domain.Task) TaskDTOResponse {
	return TaskDTOResponse{
		ID:           domain.ID,
		Version:      domain.Version,
		Title:        domain.Title,
		Description:  domain.Description,
		Completed:    domain.Completed,
		CreatedAt:    domain.CreatedAt,
		CompletedAt:  domain.CompletedAt,
		AuthorUserID: domain.AuthorUserID,
	}
}

func tasksDTOFromDomains(tasks []domain.Task) []TaskDTOResponse {
	tasksDTO := make([]TaskDTOResponse, len(tasks))
	for i, task := range tasks {
		tasksDTO[i] = taskDTOFromDomain(task)
	}

	return tasksDTO
}

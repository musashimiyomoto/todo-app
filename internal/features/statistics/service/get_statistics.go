package statistics_service

import (
	"context"
	"fmt"
	"time"

	domain "github.com/musashimiyomoto/todo-app/internal/core/domain"
	core_errors "github.com/musashimiyomoto/todo-app/internal/core/errors"
)

func (s *StatisticsService) GetStatistics(
	ctx context.Context,
	userID *int,
	from *time.Time,
	to *time.Time,
) (domain.Statistics, error) {
	if from != nil && to != nil {
		if to.Before(*from) || to.Equal(*from) {
			return domain.Statistics{}, fmt.Errorf(
				"'to' time must be after 'from': %w",
				core_errors.ErrInvalidArgument,
			)
		}
	}

	taskDomains, err := s.statisticsRepository.GetTasks(ctx, userID, from, to)
	if err != nil {
		return domain.Statistics{}, fmt.Errorf("Get tasks from repository: %w", err)
	}

	return calculateStatistics(taskDomains), nil
}

func calculateStatistics(taskDomains []domain.Task) domain.Statistics {
	tasksCreated := len(taskDomains)

	if tasksCreated == 0 {
		return domain.Statistics{}
	}

	tasksCompleted := 0
	var totalCompletionDuration time.Duration
	for _, task := range taskDomains {
		if task.Completed {
			tasksCompleted++
		}

		completionDuration := task.CompletionDuration()
		if completionDuration != nil {
			totalCompletionDuration += *completionDuration
		}
	}

	tasksCompletedRate := float64(tasksCompleted) / float64(tasksCreated) * 100

	var tasksAverageCompletionTime *time.Duration
	if tasksCompleted > 0 && totalCompletionDuration != 0 {
		avg := totalCompletionDuration / time.Duration(tasksCompleted)
		tasksAverageCompletionTime = &avg
	}

	return domain.Statistics{
		TasksCreated:               tasksCreated,
		TasksCompleted:             tasksCompleted,
		TasksCompletedRate:         &tasksCompletedRate,
		TasksAverageCompletionTime: tasksAverageCompletionTime,
	}
}

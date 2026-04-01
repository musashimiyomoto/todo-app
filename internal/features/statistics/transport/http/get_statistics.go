package statistics_transport_http

import (
	"fmt"
	"net/http"
	"time"

	core_logger "github.com/musashimiyomoto/todo-app/internal/core/core_logger"
	domain "github.com/musashimiyomoto/todo-app/internal/core/domain"
	core_http_request "github.com/musashimiyomoto/todo-app/internal/core/transport/http/request"
	core_http_response "github.com/musashimiyomoto/todo-app/internal/core/transport/http/response"
)

type GetStatisticsResponse struct {
	TasksCreated               int
	TasksCompleted             int
	TasksCompletedRate         *float64
	TasksAverageCompletionTime *string
}

func (h *StatisticsHTTPHandler) GetStatistics(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, from, to, err := getUserIDFromToQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "Failed to get 'user_id', 'from', and 'to' query params")

		return
	}

	statisticsDomain, err := h.statisticsService.GetStatistics(ctx, userID, from, to)
	if err != nil {
		responseHandler.ErrorResponse(err, "Failed to get statistics")

		return
	}

	responseHandler.JSONResponse(statisticsDTOFromDomain(statisticsDomain), http.StatusOK)
}

func statisticsDTOFromDomain(statisticsDomain domain.Statistics) GetStatisticsResponse {
	var tasksAverageCompletionTime *string

	if statisticsDomain.TasksAverageCompletionTime != nil {
		duration := statisticsDomain.TasksAverageCompletionTime.String()
		tasksAverageCompletionTime = &duration
	}

	return GetStatisticsResponse{
		TasksCreated:               statisticsDomain.TasksCreated,
		TasksCompleted:             statisticsDomain.TasksCompleted,
		TasksCompletedRate:         statisticsDomain.TasksCompletedRate,
		TasksAverageCompletionTime: tasksAverageCompletionTime,
	}
}

func getUserIDFromToQueryParams(r *http.Request) (*int, *time.Time, *time.Time, error) {
	const (
		userIDQueryParamKey = "user_id"
		fromQueryParamKey   = "from"
		toQueryParamKey     = "to"
	)

	userID, err := core_http_request.GetIntQueryParam(r, userIDQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("Failed to get 'user_id' query param: %w", err)
	}

	from, err := core_http_request.GetDateQueryParam(r, fromQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("Failed to get 'from' query param: %w", err)
	}

	to, err := core_http_request.GetDateQueryParam(r, toQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("Failed to get 'to' query param: %w", err)
	}

	return userID, from, to, nil
}

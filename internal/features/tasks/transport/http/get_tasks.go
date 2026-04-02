package tasks_transport_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/musashimiyomoto/todo-app/internal/core/core_logger"
	core_http_request "github.com/musashimiyomoto/todo-app/internal/core/transport/http/request"
	core_http_response "github.com/musashimiyomoto/todo-app/internal/core/transport/http/response"
)

type GetTasksResponse []TaskDTOResponse

// GetTasks 	 godoc
// @Summary 	 Get tasks
// @Description  Returns a list of tasks with optional filtering and pagination.
// @Tags 		 Tasks
// @Produce 	 json
// @Param 		 user_id query int false "Filter by user ID"
// @Param 		 limit   query int false "Limit"
// @Param 		 offset  query int false "Offset"
// @Success 	 200 {object} GetTasksResponse "List of tasks"
// @Failure 	 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 	 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router 		 /tasks [get]
func (h *TasksHTTPHandler) GetTasks(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, limit, offset, err := getUserIDLimitOffsetQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "Failed to get 'user_id', 'limit', and 'offset' query params")

		return
	}

	taskDomains, err := h.tasksService.GetTasks(ctx, userID, limit, offset)
	if err != nil {
		responseHandler.ErrorResponse(err, "Failed to get tasks")

		return
	}

	responseHandler.JSONResponse(
		GetTasksResponse(
			tasksDTOFromDomains(taskDomains),
		),
		http.StatusOK,
	)
}

func getUserIDLimitOffsetQueryParams(r *http.Request) (*int, *int, *int, error) {
	const (
		userIDQueryParamKey = "user_id"
		limitQueryParamKey  = "limit"
		offsetQueryParamKey = "offset"
	)

	userID, err := core_http_request.GetIntQueryParam(r, userIDQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("Failed to get 'user_id' query param: %w", err)
	}

	limit, err := core_http_request.GetIntQueryParam(r, limitQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("Failed to get 'limit' query param: %w", err)
	}

	offset, err := core_http_request.GetIntQueryParam(r, offsetQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("Failed to get 'offset' query param: %w", err)
	}

	return userID, limit, offset, nil
}

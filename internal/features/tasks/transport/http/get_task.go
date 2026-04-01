package tasks_transport_http

import (
	"net/http"

	core_logger "github.com/musashimiyomoto/todo-app/internal/core/core_logger"
	core_http_request "github.com/musashimiyomoto/todo-app/internal/core/transport/http/request"
	core_http_response "github.com/musashimiyomoto/todo-app/internal/core/transport/http/response"
)

type GetTaskResponse TaskDTOResponse

func (h *TasksHTTPHandler) GetTask(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	taskID, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "Failed to get 'id' path value")

		return
	}

	taskDomain, err := h.tasksService.GetTask(ctx, taskID)
	if err != nil {
		responseHandler.ErrorResponse(err, "Failed to get task")

		return
	}

	responseHandler.JSONResponse(
		GetTaskResponse(
			taskDTOFromDomain(taskDomain),
		),
		http.StatusOK,
	)
}

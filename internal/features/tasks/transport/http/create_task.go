package tasks_transport_http

import (
	"net/http"

	core_logger "github.com/musashimiyomoto/todo-app/internal/core/core_logger"
	domain "github.com/musashimiyomoto/todo-app/internal/core/domain"
	core_http_request "github.com/musashimiyomoto/todo-app/internal/core/transport/http/request"
	core_http_response "github.com/musashimiyomoto/todo-app/internal/core/transport/http/response"
)

type CreateTaskRequest struct {
	Title        string  `json:"title" validate:"required,min=1,max=100"`
	Description  *string `json:"description" validate:"omitempty,min=1,max=1000"`
	AuthorUserID int     `json:"author_user_id" validate:"required"`
}

type CreateTaskResponse TaskDTOResponse

// CreateTask 	 godoc
// @Summary 	 Create task
// @Description  Creates a new task in the system.
// @Tags 		 Tasks
// @Accept 		 json
// @Produce 	 json
// @Param 		 request body CreateTaskRequest true "Create task request"
// @Success 	 201 {object} CreateTaskResponse "Task created successfully"
// @Failure 	 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 	 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router 		 /tasks [post]
func (h *TasksHTTPHandler) CreateTask(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	var request CreateTaskRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "Failed to decode and validate HTTP request")

		return
	}

	taskDomain, err := h.tasksService.CreateTask(ctx, taskDomainFromDTO(request))
	if err != nil {
		responseHandler.ErrorResponse(err, "Failed to create task")

		return
	}

	responseHandler.JSONResponse(
		CreateTaskResponse(
			taskDTOFromDomain(taskDomain),
		),
		http.StatusCreated,
	)
}

func taskDomainFromDTO(dto CreateTaskRequest) domain.Task {
	return domain.NewTaskUninitialized(dto.Title, dto.Description, dto.AuthorUserID)
}

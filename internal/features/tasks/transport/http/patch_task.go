package tasks_transport_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/musashimiyomoto/todo-app/internal/core/core_logger"
	domain "github.com/musashimiyomoto/todo-app/internal/core/domain"
	core_errors "github.com/musashimiyomoto/todo-app/internal/core/errors"
	core_http_request "github.com/musashimiyomoto/todo-app/internal/core/transport/http/request"
	core_http_response "github.com/musashimiyomoto/todo-app/internal/core/transport/http/response"
	core_http_types "github.com/musashimiyomoto/todo-app/internal/core/transport/http/types"
)

type PatchTaskRequest struct {
	Title       core_http_types.Nullable[string] `json:"title"`
	Description core_http_types.Nullable[string] `json:"description"`
	Completed   core_http_types.Nullable[bool]   `json:"completed"`
}

func (r PatchTaskRequest) Validate() error {
	if r.Title.Set {
		if r.Title.Value == nil {
			return fmt.Errorf("`Title` cannot be null: %w", core_errors.ErrInvalidArgument)
		}

		titleLength := len([]rune(*r.Title.Value))
		if titleLength < 1 || titleLength > 100 {
			return fmt.Errorf("Invalid `Title` len: %d: %w", titleLength, core_errors.ErrInvalidArgument)
		}
	}

	if r.Description.Set {
		if r.Description.Value != nil {
			descriptionLength := len([]rune(*r.Description.Value))
			if descriptionLength < 1 || descriptionLength > 1000 {
				return fmt.Errorf("Invalid `Description` len: %d: %w", descriptionLength, core_errors.ErrInvalidArgument)
			}
		}
	}

	if r.Completed.Set {
		if r.Completed.Value == nil {
			return fmt.Errorf("`Completed` cannot be null: %w", core_errors.ErrInvalidArgument)
		}
	}

	return nil
}

type PatchTaskResponse TaskDTOResponse

// PatchTask 	 godoc
// @Summary 	 Patch task
// @Description  Partially updates a task by ID.
// @Tags 		 Tasks
// @Accept 		 json
// @Produce 	 json
// @Param 		 id      path int              true "Task ID"
// @Param 		 request body PatchTaskRequest  true "Patch task request"
// @Success 	 200 {object} PatchTaskResponse "Task updated successfully"
// @Failure 	 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 	 404 {object} core_http_response.ErrorResponse "Task not found"
// @Failure 	 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router 		 /tasks/{id} [patch]
func (h *TasksHTTPHandler) PatchTask(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	taskID, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "Failed to get 'id' path value")

		return
	}

	var request PatchTaskRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "Failed to decode and validate HTTP request")

		return
	}

	taskDomain, err := h.tasksService.PatchTask(ctx, taskID, taskPatchFromRequest(request))
	if err != nil {
		responseHandler.ErrorResponse(err, "Failed to patch task")

		return
	}

	responseHandler.JSONResponse(
		PatchTaskResponse(
			taskDTOFromDomain(taskDomain),
		),
		http.StatusOK,
	)
}

func taskPatchFromRequest(request PatchTaskRequest) domain.TaskPatch {
	return domain.NewTaskPatch(
		request.Title.ToDomain(),
		request.Description.ToDomain(),
		request.Completed.ToDomain(),
	)
}

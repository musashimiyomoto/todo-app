package users_transport_http

import (
	"net/http"

	core_logger "github.com/musashimiyomoto/todo-app/internal/core/core_logger"
	core_http_request "github.com/musashimiyomoto/todo-app/internal/core/transport/http/request"
	core_http_response "github.com/musashimiyomoto/todo-app/internal/core/transport/http/response"
)

type GetUserResponse UserDTOResponse

// GetUser 		 godoc
// @Summary 	 Get user by ID
// @Description  Returns a single user by ID.
// @Tags 		 Users
// @Produce 	 json
// @Param 		 id path int true "User ID"
// @Success 	 200 {object} GetUserResponse "User found"
// @Failure 	 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 	 404 {object} core_http_response.ErrorResponse "User not found"
// @Failure 	 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router 		 /users/{id} [get]
func (h *UsersHTTPHandler) GetUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "Failed to get 'id' path value")

		return
	}

	userDomain, err := h.usersService.GetUser(ctx, userID)
	if err != nil {
		responseHandler.ErrorResponse(err, "Failed to get user")

		return
	}

	responseHandler.JSONResponse(
		GetUserResponse(
			userDTOFromDomain(userDomain),
		),
		http.StatusOK,
	)
}

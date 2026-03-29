package users_transport_http

import (
	"net/http"

	core_loger "github.com/musashimiyomoto/todo-app/internal/core/logger"
	core_http_response "github.com/musashimiyomoto/todo-app/internal/core/transport/http/response"
	core_http_utils "github.com/musashimiyomoto/todo-app/internal/core/transport/http/utils"
)

type GetUserResponse UserDTOResponse

func (h *UsersHTTPHandler) GetUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_loger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, err := core_http_utils.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "Failed to get 'id' path value")

		return
	}

	userDomain, err := h.usersService.GetUser(ctx, userID)
	if err != nil {
		responseHandler.ErrorResponse(err, "Failed to get user")

		return
	}

	responseHandler.JSONResponse(GetUserResponse(userDTOFromDomain(userDomain)), http.StatusOK)
}

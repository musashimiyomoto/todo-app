package users_transport_http

import (
	"fmt"
	"net/http"

	core_loger "github.com/musashimiyomoto/todo-app/internal/core/logger"
	core_http_response "github.com/musashimiyomoto/todo-app/internal/core/transport/http/response"
	core_http_utils "github.com/musashimiyomoto/todo-app/internal/core/transport/http/utils"
)

type GetUsersResponse []UserDTOResponse

func (h *UsersHTTPHandler) GetUsers(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_loger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	limit, offset, err := getLimitOffsetQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "Failed to get 'limit' and 'offset' query params")

		return
	}

	userDomains, err := h.usersService.GetUsers(ctx, limit, offset)
	if err != nil {
		responseHandler.ErrorResponse(err, "Failed to get users")

		return
	}

	responseHandler.JSONResponse(GetUsersResponse(usersDTOFromDomains(userDomains)), http.StatusOK)
}

func getLimitOffsetQueryParams(r *http.Request) (*int, *int, error) {
	limit, err := core_http_utils.GetIntQueryParam(r, "limit")
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to get 'limit' query param: %w", err)
	}

	offset, err := core_http_utils.GetIntQueryParam(r, "offset")
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to get 'offset' query param: %w", err)
	}

	return limit, offset, nil
}

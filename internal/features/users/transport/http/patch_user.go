package users_transport_http

import (
	"fmt"
	"net/http"
	"strings"

	core_logger "github.com/musashimiyomoto/todo-app/internal/core/core_logger"
	domain "github.com/musashimiyomoto/todo-app/internal/core/domain"
	core_errors "github.com/musashimiyomoto/todo-app/internal/core/errors"
	core_http_request "github.com/musashimiyomoto/todo-app/internal/core/transport/http/request"
	core_http_response "github.com/musashimiyomoto/todo-app/internal/core/transport/http/response"
	core_http_types "github.com/musashimiyomoto/todo-app/internal/core/transport/http/types"
)

type PatchUserRequest struct {
	FullName    core_http_types.Nullable[string] `json:"full_name"`
	PhoneNumber core_http_types.Nullable[string] `json:"phone_number"`
}

func (r PatchUserRequest) Validate() error {
	if r.FullName.Set {
		if r.FullName.Value == nil {
			return fmt.Errorf("`FullName` cannot be null: %w", core_errors.ErrInvalidArgument)
		}

		fullNameLength := len([]rune(*r.FullName.Value))
		if fullNameLength < 3 || fullNameLength > 100 {
			return fmt.Errorf("Invalid `FullName` len: %d: %w", fullNameLength, core_errors.ErrInvalidArgument)
		}
	}

	if r.PhoneNumber.Set {
		if r.PhoneNumber.Value != nil {
			phoneNumberLength := len([]rune(*r.PhoneNumber.Value))
			if phoneNumberLength < 10 || phoneNumberLength > 15 {
				return fmt.Errorf("Invalid `PhoneNumber` len: %d: %w", phoneNumberLength, core_errors.ErrInvalidArgument)
			}

			if !strings.HasPrefix(*r.PhoneNumber.Value, "+") {
				return fmt.Errorf("`PhoneNumber` must starts with +: %w", core_errors.ErrInvalidArgument)
			}
		}
	}

	return nil
}

type PatchUserResponse UserDTOResponse

func (h *UsersHTTPHandler) PatchUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	userID, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "Failed to get 'id' path value")

		return
	}

	var request PatchUserRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "Failed to decode and validate HTTP request")

		return
	}

	userDomain, err := h.usersService.PatchUser(ctx, userID, userPatchFromRequest(request))
	if err != nil {
		responseHandler.ErrorResponse(err, "Failed to patch user")

		return
	}

	responseHandler.JSONResponse(
		PatchUserResponse(
			userDTOFromDomain(userDomain),
		),
		http.StatusOK,
	)
}

func userPatchFromRequest(request PatchUserRequest) domain.UserPatch {
	return domain.NewUserPatch(
		request.FullName.ToDomain(),
		request.PhoneNumber.ToDomain(),
	)
}

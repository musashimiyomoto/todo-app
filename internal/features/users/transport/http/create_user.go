package users_transport_http

import (
	"net/http"

	"github.com/musashimiyomoto/todo-app/internal/core/domain"
	core_loger "github.com/musashimiyomoto/todo-app/internal/core/logger"
	core_http_request "github.com/musashimiyomoto/todo-app/internal/core/transport/http/request"
	core_http_response "github.com/musashimiyomoto/todo-app/internal/core/transport/http/response"
)

type CreateUserRequest struct {
	FullName    string  `json:"full_name" validate:"required,min=3,max=100"`
	PhoneNumber *string `json:"phone_number" validate:"omitempty,min=10,max=15,startsWith=+"`
}

type CreateUserResponse struct {
	ID          int     `json:"id"`
	Version     int     `json:"version"`
	FullName    string  `json:"full_name"`
	PhoneNumber *string `json:"phone_number"`
}

func (h *UsersHTTPHandler) CreateUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_loger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	log.Debug("Create user handler")

	var request CreateUserRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "Failed to decode and validate HTTP request")

		return
	}

	user, err := h.usersService.CreateUser(ctx, domainFromDTO(request))
	if err != nil {
		responseHandler.ErrorResponse(err, "Failed to create user")

		return
	}

	responseHandler.JSONResponse(dtoFromDomain(user), http.StatusCreated)
}

func domainFromDTO(dto CreateUserRequest) domain.User {
	return domain.NewUserUninitialized(dto.FullName, dto.PhoneNumber)
}

func dtoFromDomain(domain domain.User) CreateUserResponse {
	return CreateUserResponse{
		ID:          domain.ID,
		Version:     domain.Version,
		FullName:    domain.FullName,
		PhoneNumber: domain.PhoneNumber,
	}
}

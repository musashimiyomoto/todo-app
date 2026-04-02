package users_transport_http

import (
	"net/http"

	core_logger "github.com/musashimiyomoto/todo-app/internal/core/core_logger"
	domain "github.com/musashimiyomoto/todo-app/internal/core/domain"
	core_http_request "github.com/musashimiyomoto/todo-app/internal/core/transport/http/request"
	core_http_response "github.com/musashimiyomoto/todo-app/internal/core/transport/http/response"
)

type CreateUserRequest struct {
	FullName    string  `json:"full_name" validate:"required,min=3,max=100"`
	PhoneNumber *string `json:"phone_number" validate:"omitempty,min=10,max=15,startswith=+"`
}

type CreateUserResponse UserDTOResponse

// CreateUser 	 godoc
// @Summary 	 Create user
// @Description  Creates a new user in the system.
// @Tags 		 Users
// @Accept 		 json
// @Produce 	 json
// @Param 		 request body CreateUserRequest true "Create user request"
// @Success 	 201 {object} CreateUserResponse "User created successfully"
// @Failure 	 400 {object} core_http_response.ErrorResponse "Bad request"
// @Failure 	 500 {object} core_http_response.ErrorResponse "Internal server error"
// @Router 		 /users [post]
func (h *UsersHTTPHandler) CreateUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, rw)

	var request CreateUserRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "Failed to decode and validate HTTP request")

		return
	}

	userDomain, err := h.usersService.CreateUser(ctx, userDomainFromDTO(request))
	if err != nil {
		responseHandler.ErrorResponse(err, "Failed to create user")

		return
	}

	responseHandler.JSONResponse(
		CreateUserResponse(
			userDTOFromDomain(userDomain),
		),
		http.StatusCreated,
	)
}

func userDomainFromDTO(dto CreateUserRequest) domain.User {
	return domain.NewUserUninitialized(dto.FullName, dto.PhoneNumber)
}

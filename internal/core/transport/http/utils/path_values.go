package core_http_utils

import (
	"fmt"
	"net/http"
	"strconv"

	core_errors "github.com/musashimiyomoto/todo-app/internal/core/errors"
)

func GetIntPathValue(r *http.Request, key string) (int, error) {
	pathValue := r.PathValue(key)

	if pathValue == "" {
		return 0, fmt.Errorf(
			"Path value for key='%s' not found: %w",
			key, core_errors.ErrInvalidArgument,
		)
	}

	val, err := strconv.Atoi(pathValue)
	if err != nil {
		return 0, fmt.Errorf(
			"Failed to parse path value='%s' for key='%s': %v: %w",
			pathValue, key, err, core_errors.ErrInvalidArgument,
		)
	}

	return val, nil
}

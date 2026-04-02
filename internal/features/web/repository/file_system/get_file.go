package web_fs_repository

import (
	"errors"
	"fmt"
	"os"

	core_errors "github.com/musashimiyomoto/todo-app/internal/core/errors"
)

func (r *WebRepository) GetMainPage(filePath string) ([]byte, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, fmt.Errorf("File:%s: %w", filePath, core_errors.ErrNotFound)
		}

		return nil, fmt.Errorf("Get file: %s: %w", filePath, err)
	}

	return file, nil
}

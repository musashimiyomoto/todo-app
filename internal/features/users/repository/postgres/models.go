package users_postgres_repository

import domain "github.com/musashimiyomoto/todo-app/internal/core/domain"

type UserModel struct {
	ID          int
	Version     int
	FullName    string
	PhoneNumber *string
}

func userDomainsFromModels(userModels []UserModel) []domain.User {
	usersDomains := make([]domain.User, len(userModels))

	for i, userModel := range userModels {
		usersDomains[i] = domain.NewUser(
			userModel.ID,
			userModel.Version,
			userModel.FullName,
			userModel.PhoneNumber,
		)
	}

	return usersDomains
}

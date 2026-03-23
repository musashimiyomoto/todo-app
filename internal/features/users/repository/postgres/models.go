package users_postgres_repository

type UserModel struct {
	ID          int     `db:"id"`
	Version     int     `db:"version"`
	FullName    string  `db:"full_name"`
	PhoneNumber *string `db:"phone_number"`
}

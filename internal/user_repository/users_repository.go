package user_repository

type UserCore struct {
	UserId string
}

type UserRepository interface {
	GetUserByID(id string) (UserCore, error)
	GetUsers() ([]UserCore, error)
	CreateUser(user UserCore) error
	DeleteUser(id string) error
}

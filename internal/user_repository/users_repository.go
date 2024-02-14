package user_repository

import "errors"

type UserCore struct {
	UserId string
}

var ErrUserIDNotFound = errors.New("userId not found")

type UserRepository interface {
	GetUserByID(id string) (*UserCore, error)
	GetUsers() ([]UserCore, error)
	CreateUser(user UserCore) error
	DeleteUser(id string) error
}

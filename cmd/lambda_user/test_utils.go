package main

import "github.com/safetorun/PromptDefender/user_repository"

type MockUserRepository struct {
	getUsersResponse   []user_repository.UserCore
	onCreateUserCalled *func()
}

func (m MockUserRepository) GetUsers(s string) ([]user_repository.UserCore, error) {
	return m.getUsersResponse, nil
}

func (m MockUserRepository) CreateUser(user user_repository.UserCore) error {
	if m.onCreateUserCalled != nil {
		(*m.onCreateUserCalled)()
	}
	return nil
}

func (m MockUserRepository) DeleteUser(id string, apikey string) error {
	//TODO implement me
	panic("implement me")
}

func (m MockUserRepository) GetUserByID(id string, apikey string) (*user_repository.UserCore, error) {
	return &user_repository.UserCore{UserOrSessionId: id}, nil
}

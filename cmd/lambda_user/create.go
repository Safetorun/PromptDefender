package main

import (
	"github.com/safetorun/PromptDefender/user_repository"
	"github.com/safetorun/PromptDefender/user_repository_ddb"
)

type CreateUserHandler struct {
	userInstance user_repository.UserRepository
	apiKey       string
}

func NewCreateUserHandler(apiKey string) *CreateUserHandler {
	return &CreateUserHandler{
		userInstance: user_repository_ddb.New(),
		apiKey:       apiKey,
	}
}

func (h *CreateUserHandler) Handle(user User) (*User, error) {
	err := h.userInstance.CreateUser(user_repository.UserCore{
		UserOrSessionId: *user.UserId,
		ApiKeyId:        h.apiKey,
	})

	if err != nil {
		return nil, err
	}

	return &user, nil
}

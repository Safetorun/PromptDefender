package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/safetorun/PromptDefender/user_repository"
	"github.com/safetorun/PromptDefender/user_repository_ddb"
	"github.com/thoas/go-funk"
	"log"
)

type RetrieveUserHandler struct {
	userInstance user_repository.UserRepository
	apikeyId     string
	logger       *log.Logger
}

func NewRetrieveHandler(apiKeyId string) *RetrieveUserHandler {
	return &RetrieveUserHandler{
		userInstance: user_repository_ddb.New(),
		logger:       log.Default(),
		apikeyId:     apiKeyId,
	}
}

func MapUsersToUserCores(users []user_repository.UserCore) []User {
	userCores := funk.Map(users, func(u user_repository.UserCore) User {
		return User{UserId: &u.UserOrSessionId}
	}).([]User)

	return userCores
}

func (h *RetrieveUserHandler) Handle() events.APIGatewayProxyResponse {
	users, err := h.userInstance.GetUsers(h.apikeyId)

	if err != nil {
		h.logger.Println(fmt.Printf("An error occurred while retrieving users: %s", err))
		return events.APIGatewayProxyResponse{StatusCode: 400}
	}

	result, err := json.Marshal(MapUsersToUserCores(users))

	if err != nil {
		h.logger.Println(fmt.Printf("An error occurred while marhsalling response: %s", err))
		return events.APIGatewayProxyResponse{StatusCode: 400}
	}

	return events.APIGatewayProxyResponse{StatusCode: 200, Body: string(result)}
}

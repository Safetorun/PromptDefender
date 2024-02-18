package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/safetorun/PromptDefender/user_repository"
	"github.com/safetorun/PromptDefender/user_repository_ddb"
	"log"
)

type DeleteUserHandler struct {
	userInstance user_repository.UserRepository
	logger       *log.Logger
	apiKeyId     string
}

func NewDeleteUserHandler(apiKeyId string) *DeleteUserHandler {
	return &DeleteUserHandler{
		userInstance: user_repository_ddb.New(),
		logger:       log.Default(),
		apiKeyId:     apiKeyId,
	}
}

func (h *DeleteUserHandler) Handle(userId string) events.APIGatewayProxyResponse {
	h.logger.Println("Deleting user with it of: ", userId)

	err := h.userInstance.DeleteUser(userId, h.apiKeyId)

	if err != nil {
		h.logger.Println("Error deleting user: ", err)
		return events.APIGatewayProxyResponse{StatusCode: 400}
	}

	return events.APIGatewayProxyResponse{StatusCode: 204}
}

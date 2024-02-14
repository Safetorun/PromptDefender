package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/safetorun/PromptDefender/user_repository_ddb"
)

type DeleteUserHandler struct {
	userInstance *user_repository_ddb.UserRepositoryDdb
}

func (h *DeleteUserHandler) Handle(userId string) events.APIGatewayProxyResponse {
	err := h.userInstance.DeleteUser(userId)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400}
	}

	return events.APIGatewayProxyResponse{StatusCode: 201}
}

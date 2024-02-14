package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/safetorun/PromptDefender/user_repository_ddb"
)

type RetrieveUserHandler struct {
	userInstance *user_repository_ddb.UserRepositoryDdb
}

func (h *RetrieveUserHandler) Handle() events.APIGatewayProxyResponse {
	users, err := h.userInstance.GetUsers()
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400}
	}

	result, err := json.Marshal(users)

	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400}
	}

	return events.APIGatewayProxyResponse{StatusCode: 200, Body: string(result)}
}

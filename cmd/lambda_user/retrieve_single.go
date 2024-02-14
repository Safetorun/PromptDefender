package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/safetorun/PromptDefender/user_repository_ddb"
)

type RetrieveUserHandlerSingle struct {
	userInstance *user_repository_ddb.UserRepositoryDdb
}

func (h *RetrieveUserHandlerSingle) Handle(userId string) events.APIGatewayProxyResponse {
	user, err := h.userInstance.GetUserByID(userId)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400}
	}

	result, err := json.Marshal(user)

	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400}
	}

	return events.APIGatewayProxyResponse{StatusCode: 200, Body: string(result)}
}

package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/safetorun/PromptDefender/user_repository_ddb"
	"log"
)

type RetrieveUserHandler struct {
	userInstance *user_repository_ddb.UserRepositoryDdb
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

func (h *RetrieveUserHandler) Handle() events.APIGatewayProxyResponse {
	users, err := h.userInstance.GetUsers(h.apikeyId)
	if err != nil {
		h.logger.Println(fmt.Printf("An error occurred while retrieving users: %s", err))
		return events.APIGatewayProxyResponse{StatusCode: 400}
	}

	result, err := json.Marshal(users)

	if err != nil {
		h.logger.Println(fmt.Printf("An error occurred while marhsalling response: %s", err))
		return events.APIGatewayProxyResponse{StatusCode: 400}
	}

	return events.APIGatewayProxyResponse{StatusCode: 200, Body: string(result)}
}

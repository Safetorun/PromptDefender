package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/safetorun/PromptDefender/user_repository"
	"github.com/safetorun/PromptDefender/user_repository_ddb"
	"log"
)

type RetrieveUserHandlerSingle struct {
	userInstance user_repository.UserRepository
	logger       *log.Logger
	apiKey       string
}

func NewRetrieverHandlerSingle(apiKey string) *RetrieveUserHandlerSingle {
	return &RetrieveUserHandlerSingle{
		userInstance: user_repository_ddb.New(),
		logger:       log.Default(),
		apiKey:       apiKey,
	}
}

func (h *RetrieveUserHandlerSingle) Handle(userId string) events.APIGatewayProxyResponse {
	user, err := h.userInstance.GetUserByID(userId, h.apiKey)

	if errors.Is(err, user_repository.ErrUserIDNotFound) {
		h.logger.Println("User not found with id ", userId)
		return events.APIGatewayProxyResponse{StatusCode: 404}
	}

	if err != nil {
		h.logger.Println(fmt.Sprintf("Error retrieving user: %v", err))
		return events.APIGatewayProxyResponse{StatusCode: 400}
	}

	userResponse := User{
		UserId: &user.UserOrSessionId,
	}

	result, err := json.Marshal(userResponse)

	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400}
	}

	return events.APIGatewayProxyResponse{StatusCode: 200, Body: string(result)}
}

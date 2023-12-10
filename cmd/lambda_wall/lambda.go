package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/safetorun/PromptDefender/aiprompt"
	"github.com/safetorun/PromptDefender/aws/base_aws"
	"github.com/safetorun/PromptDefender/wall"
	"os"
)

type WallLambda struct {
	wallInstance *wall.Wall
}

func (w *WallLambda) Handle(promptRequest WallRequest) (*WallResponse, error) {
	answer, err := w.wallInstance.CheckWall(wall.PromptToCheck{Prompt: promptRequest.Prompt})

	if err != nil {
		return nil, err
	}

	return &WallResponse{InjectionScore: &answer.Score}, nil
}

func Handler(_ context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	openAIKey, exists := os.LookupEnv("open_ai_api_key")

	if !exists {
		return events.APIGatewayProxyResponse{StatusCode: 400}, fmt.Errorf("error retrieving API key: environment variable not set")
	}

	wallBuilder := wall.New(aiprompt.NewOpenAI(openAIKey))

	handler := WallLambda{wallInstance: &wallBuilder}

	response, err := base_aws.BaseHandler[WallRequest, WallResponse](request, &handler)

	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400}, err
	}

	return response, nil
}

func main() {
	lambda.Start(Handler)
}

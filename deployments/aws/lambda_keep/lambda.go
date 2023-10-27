package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/safetorun/PromptDefender/aiprompt"
	"github.com/safetorun/PromptDefender/aws/base_aws"
	"github.com/safetorun/PromptDefender/keep"
	"os"
)

type PromptBuilderResponse struct {
	NewPrompt string
}

type PromptBuilderRequest struct {
	Prompt string `json:"prompt"`
}

type KeepLambda struct {
	keepInstance *keep.Keep
}

func (k *KeepLambda) Handle(promptRequest KeepRequest) (*KeepResponse, error) {
	answer, err := k.keepInstance.BuildKeep(keep.StartingPrompt{Prompt: promptRequest.Prompt})

	if err != nil {
		return nil, err
	}

	return &KeepResponse{ShieldedPrompt: &answer.NewPrompt}, nil
}

func Handler(_ context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	openAIKey, exists := os.LookupEnv("open_ai_api_key")

	if !exists {
		return events.APIGatewayProxyResponse{StatusCode: 400}, fmt.Errorf("error retrieving API key: environment variable not set")
	}

	keepBuilder := keep.New(aiprompt.NewOpenAI(openAIKey))
	handler := KeepLambda{keepInstance: keepBuilder}

	response, err := base_aws.BaseHandler[KeepRequest, KeepResponse](request, &handler)

	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400}, err
	}

	return response, nil
}

func main() {
	lambda.Start(Handler)
}

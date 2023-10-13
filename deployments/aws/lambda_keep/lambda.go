package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/safetorun/PromptShield/aiprompt"
	"github.com/safetorun/PromptShield/app"
)

type PromptBuilderResponse struct {
	NewPrompt string
}

type PromptBuilderRequest struct {
	Prompt string `json:"prompt"`
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	openAIKey, exists := os.LookupEnv("open_ai_api_key")

	if !exists {
		return events.APIGatewayProxyResponse{StatusCode: 400}, fmt.Errorf("error retrieving API key: environment variable not set")
	}

	var promptRequest PromptBuilderRequest

	if err := json.Unmarshal([]byte(request.Body), &promptRequest); err != nil {
		fmt.Printf("error unmarshalling request: %v\n", err)
		return events.APIGatewayProxyResponse{StatusCode: 400}, err
	}

	fmt.Printf("Received request for %v\n", promptRequest)

	answer, err := app.New(*aiprompt.NewOpenAI(openAIKey)).BuildPromptDefense(promptRequest.Prompt)

	if err != nil {
		fmt.Printf("error processing AI: %v\n", err)
		return events.APIGatewayProxyResponse{StatusCode: 400}, fmt.Errorf("error processing AI request: %v", err)
	}

	response := PromptBuilderResponse{NewPrompt: *answer}

	jsonBytes, err := json.Marshal(response)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400}, err
	}

	jsonString := string(jsonBytes)

	return events.APIGatewayProxyResponse{StatusCode: 200, Body: jsonString}, nil
}

func main() {
	lambda.Start(Handler)
}

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

type AppResponse struct {
	AiScore float32 `json:"injection_score"`
}

type PromptRequest struct {
	Prompt string `json:"prompt"`
}

type APIGatewayResponse struct {
	StatusCode        int                 `json:"statusCode"`
	Headers           map[string]string   `json:"headers,omitempty"`
	MultiValueHeaders map[string][]string `json:"multiValueHeaders,omitempty"`
	Body              string              `json:"body"`
	IsBase64Encoded   bool                `json:"isBase64Encoded,omitempty"`
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	openAIKey, exists := os.LookupEnv("open_ai_api_key")

	if !exists {
		return events.APIGatewayProxyResponse{StatusCode: 400}, fmt.Errorf("error retrieving API key: environment variable not set")
	}

	var promptRequest PromptRequest

	if err := json.Unmarshal([]byte(request.Body), &promptRequest); err != nil {
		// Handle error
		return events.APIGatewayProxyResponse{StatusCode: 400}, err
	}

	answer, err := app.New(*aiprompt.NewOpenAI(openAIKey)).CheckAI(promptRequest.Prompt)

	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400}, fmt.Errorf("error processing AI request: %v", err)
	}

	response := AppResponse{AiScore: answer}

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

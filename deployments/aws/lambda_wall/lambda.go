package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/safetorun/PromptDefender/aiprompt"
	"github.com/safetorun/PromptDefender/wall"
	"os"
)

type LambdaWallResponse struct {
	AiScore float32 `json:"injection_score"`
}

type LambdaWallRequest struct {
	Prompt string `json:"prompt"`
}

func Handler(_ context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	openAIKey, exists := os.LookupEnv("open_ai_api_key")

	if !exists {
		return events.APIGatewayProxyResponse{StatusCode: 400}, fmt.Errorf("error retrieving API key: environment variable not set")
	}

	var promptRequest LambdaWallRequest

	if err := json.Unmarshal([]byte(request.Body), &promptRequest); err != nil {
		fmt.Printf("error unmarshalling request: %v\n", err)
		return events.APIGatewayProxyResponse{StatusCode: 400}, err
	}

	fmt.Printf("Received request for %v\n", promptRequest)

	wallBuilder := wall.New(aiprompt.NewOpenAI(openAIKey))
	answer, err := wallBuilder.CheckWall(wall.PromptToCheck{Prompt: promptRequest.Prompt})

	if err != nil {
		fmt.Printf("error processing AI: %v\n", err)
		return events.APIGatewayProxyResponse{StatusCode: 400}, fmt.Errorf("error processing AI request: %v", err)
	}

	response := LambdaWallResponse{AiScore: answer.Score}

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

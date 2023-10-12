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
	"github.com/safetorun/PromptShield/pii_aws"
)

type AppResponse struct {
	AiScore     float32 `json:"injection_score"`
	ContainsPii bool    `json:"contains_pii"`
}

type PromptRequest struct {
	Prompt  string `json:"prompt"`
	ScanPii bool   `json:"scan_pii"`
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	openAIKey, exists := os.LookupEnv("open_ai_api_key")

	if !exists {
		return events.APIGatewayProxyResponse{StatusCode: 400}, fmt.Errorf("error retrieving API key: environment variable not set")
	}

	var promptRequest PromptRequest

	if err := json.Unmarshal([]byte(request.Body), &promptRequest); err != nil {
		fmt.Printf("error unmarshalling request: %v\n", err)
		return events.APIGatewayProxyResponse{StatusCode: 400}, err
	}

	fmt.Printf("Received request for %v\n", promptRequest)

	answer, err := app.New(*aiprompt.NewOpenAI(openAIKey)).CheckAI(promptRequest.Prompt)

	if err != nil {
		fmt.Printf("error processing AI: %v\n", err)
		return events.APIGatewayProxyResponse{StatusCode: 400}, fmt.Errorf("error processing AI request: %v", err)
	}

	var containsPii = false

	if promptRequest.ScanPii {
		piiResult, err := pii_aws.New().Scan(promptRequest.Prompt)
		if err != nil {
			fmt.Printf("Error scanning for PII %v\n", err)
			return events.APIGatewayProxyResponse{StatusCode: 400}, fmt.Errorf("error scanning for PII: %v", err)
		} else {
			containsPii = piiResult.ContainingPii
		}

	}

	response := AppResponse{AiScore: answer, ContainsPii: containsPii}

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

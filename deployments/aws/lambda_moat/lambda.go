package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/safetorun/PromptDefender/aiprompt"
	"github.com/safetorun/PromptDefender/moat"
	"github.com/safetorun/PromptDefender/pii_aws"
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
		return events.APIGatewayProxyResponse{StatusCode: 400},
			fmt.Errorf("error retrieving API key: environment variable not set")
	}

	var promptRequest PromptRequest

	if err := json.Unmarshal([]byte(request.Body), &promptRequest); err != nil {
		fmt.Printf("error unmarshalling request: %v\n", err)
		return events.APIGatewayProxyResponse{StatusCode: 400}, err
	}

	fmt.Printf("Received request for %v\n", promptRequest)

	moatInstance := moat.New(aiprompt.NewOpenAI(openAIKey), pii_aws.New())
	answer, err := moatInstance.CheckMoat(moat.PromptToCheck{
		Prompt:  promptRequest.Prompt,
		ScanPii: promptRequest.ScanPii,
	},
	)

	if err != nil {
		fmt.Printf("error processing AI: %v\n", err)
		return events.APIGatewayProxyResponse{StatusCode: 400}, fmt.Errorf("error processing AI request: %v", err)
	}

	containsPii := false
	if answer != nil && answer.PiiResult != nil {
		containsPii = answer.PiiResult.ContainsPii
	}

	response := AppResponse{AiScore: answer.InjectionScore.Score, ContainsPii: containsPii}

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

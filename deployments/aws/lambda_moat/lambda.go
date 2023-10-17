package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/safetorun/PromptDefender/moat"
	"github.com/safetorun/PromptDefender/pii_aws"
)

func Handler(_ context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var promptRequest MoatRequest

	if err := json.Unmarshal([]byte(request.Body), &promptRequest); err != nil {
		fmt.Printf("error unmarshalling request: %v\n", err)
		return events.APIGatewayProxyResponse{StatusCode: 400}, err
	}

	fmt.Printf("Received request for %v\n", promptRequest)

	moatInstance := moat.New(pii_aws.New())
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

	response := MoatResponse{ContainsPii: &containsPii}

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

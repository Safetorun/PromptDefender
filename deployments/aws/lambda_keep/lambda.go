package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/safetorun/PromptDefender/aiprompt"
	"github.com/safetorun/PromptDefender/aws/base_aws"
	"github.com/safetorun/PromptDefender/keep"
	"log"
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

var sqsQueueCallback keep.Callback = func(prompt string, newPrompt string) error {
	cfg, err := config.LoadDefaultConfig(context.Background())

	if err != nil {
		return err
	}

	queueMessage := struct {
		Request   KeepRequest
		Rersponse KeepResponse
	}{
		Request:   KeepRequest{Prompt: prompt},
		Rersponse: KeepResponse{ShieldedPrompt: &newPrompt},
	}

	jsonMessage, err := json.Marshal(queueMessage)

	if err != nil {
		log.Fatalf("JSON marshaling error: %v", err)
	}

	svc := sqs.NewFromConfig(cfg)

	queueName := os.Getenv("keep_sqs_queue_url")

	input := &sqs.SendMessageInput{
		MessageBody: aws.String(string(jsonMessage)),
		QueueUrl:    &queueName,
	}

	result, err := svc.SendMessage(context.TODO(), input)
	if err != nil {
		log.Fatalf("Error sending message: %v", err)
	}

	fmt.Printf("Message sent: %s\n", *result.MessageId)

	return nil
}

func Handler(_ context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	openAIKey, exists := os.LookupEnv("open_ai_api_key")

	if !exists {
		return events.APIGatewayProxyResponse{StatusCode: 400}, fmt.Errorf("error retrieving API key: environment variable not set")
	}

	addCallback := func(k *keep.Keep) {
		k.Callback = &sqsQueueCallback
	}

	keepBuilder := keep.New(aiprompt.NewOpenAI(openAIKey), addCallback)

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

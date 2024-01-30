// / This file contains the lambda handler for the keep lambda
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
	"github.com/safetorun/PromptDefender/internal/base_aws"
	"github.com/safetorun/PromptDefender/keep"
	"log"
	"os"
	"time"
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
	randomiseXmlTag := false
	if promptRequest.RandomiseXmlTag == nil {
		randomiseXmlTag = false
	} else {
		randomiseXmlTag = *promptRequest.RandomiseXmlTag
	}

	answer, err := k.keepInstance.BuildKeep(keep.StartingPrompt{
		Prompt:       promptRequest.Prompt,
		RandomiseTag: randomiseXmlTag,
	})

	if err != nil {
		if keep.IsPromptRequiredError(err) {
			return nil, fmt.Errorf("prompt cannot be empty")
		}

		return nil, err
	}

	return &KeepResponse{ShieldedPrompt: answer.NewPrompt, XmlTag: answer.Tag}, nil
}

var sqsQueueCallback = func(prompt string, newPrompt string, userId string, version string, request events.APIGatewayProxyRequest, startTime time.Time) error {
	cfg, err := config.LoadDefaultConfig(context.Background())

	if err != nil {
		return err
	}

	queueMessage := struct {
		Request             KeepRequest
		Response            KeepResponse
		UserId              string
		Version             string
		Endpoint            string
		Domain              string
		Headers             map[string]string
		Method              string
		QueryParams         map[string]string
		HttpMethod          string
		HttpResponse        int
		HttpResponseHeaders map[string]string
		StartedDateTime     string
		Time                int
	}{
		Endpoint:        "/keep",
		UserId:          userId,
		Version:         version,
		Request:         KeepRequest{Prompt: prompt},
		Response:        KeepResponse{ShieldedPrompt: newPrompt},
		Domain:          request.RequestContext.DomainName,
		Headers:         request.Headers,
		Method:          request.RequestContext.HTTPMethod,
		QueryParams:     request.QueryStringParameters,
		HttpMethod:      request.RequestContext.HTTPMethod,
		HttpResponse:    200,
		StartedDateTime: request.RequestContext.RequestTime,
		HttpResponseHeaders: map[string]string{
			"content-type": "application/json",
		},
		Time: int(time.Since(startTime).Milliseconds()),
	}

	jsonMessage, err := json.Marshal(queueMessage)

	if err != nil {
		log.Fatalf("JSON marshaling error: %v", err)
	}

	svc := sqs.NewFromConfig(cfg)

	queueName := retrieveQueueNameOrPanic()

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

func retrieveQueueNameOrPanic() string {
	queueName, exists := os.LookupEnv("keep_sqs_queue_url")

	if !exists {
		panic(fmt.Errorf("error retrieving API key: environment (keep_sqs_queue_url) variable not set"))
	}

	return queueName
}

// Handler is the lambda handler for the keep lambda
// The following enivorment variables are required:
// open_ai_api_key: The OpenAI API key
// version: The version of the lambda
// keep_sqs_queue_url: The SQS queue URL to send the message to
func Handler(_ context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	startTime := time.Now()
	openAIKey := retrieveApiKeyOrPanic()
	version := retrieveVersionOrPanic()
	_ = retrieveQueueNameOrPanic() // Fail early if queue name does not exist

	var addCallbackWithUserId keep.Callback = func(prompt string, newPrompt string) error {
		return sqsQueueCallback(prompt, newPrompt, request.RequestContext.Identity.APIKeyID, version, request, startTime)
	}

	addCallback := func(k *keep.Keep) {
		k.Callback = &addCallbackWithUserId
	}

	keepBuilder := keep.New(aiprompt.NewOpenAI(openAIKey), addCallback)

	handler := KeepLambda{keepInstance: keepBuilder}

	response, err := base_aws.BaseHandler[KeepRequest, KeepResponse](request, &handler)

	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400}, err
	}

	return response, nil
}

func retrieveVersionOrPanic() string {
	version, exists := os.LookupEnv("version")

	if !exists {
		panic(fmt.Errorf("error retrieving API key: environment (version) variable not set"))
	}
	return version
}

func retrieveApiKeyOrPanic() string {
	openAIKey, exists := os.LookupEnv("open_ai_api_key")

	if !exists {
		panic(fmt.Errorf("error retrieving API key: environment (openAI key) variable not set"))
	}
	return openAIKey
}

func main() {
	lambda.Start(Handler)
}

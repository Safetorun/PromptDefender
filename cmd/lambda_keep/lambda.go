// / This file contains the lambda handler for the keep lambda
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/safetorun/PromptDefender/aiprompt"
	"github.com/safetorun/PromptDefender/internal/base_aws"
	"github.com/safetorun/PromptDefender/keep"
	"go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-lambda-go/otellambda"
	"go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-lambda-go/otellambda/xrayconfig"
	"go.opentelemetry.io/contrib/propagators/aws/xray"
	"go.opentelemetry.io/otel"
	"os"
	"time"
)

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

var requestCompleteCallback = func(prompt string, newPrompt string, version string, request events.APIGatewayProxyRequest, startTime time.Time) error {

	keepRequest, err := json.Marshal(KeepRequest{Prompt: prompt})

	if err != nil {
		return err
	}

	response, err := json.Marshal(KeepResponse{ShieldedPrompt: newPrompt})

	if err != nil {
		return err
	}

	queueMessage := base_aws.ToRequestLog(request)

	queueMessage.Endpoint = "/keep"
	queueMessage.Version = version
	queueMessage.Request = string(keepRequest)
	queueMessage.Response = string(response)
	queueMessage.Time = int(time.Since(startTime).Milliseconds())

	base_aws.LogSummaryMessage(queueMessage)

	return nil
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

	var addCallbackWithUserId keep.Callback = func(prompt string, newPrompt string) error {
		return requestCompleteCallback(prompt, newPrompt, version, request, startTime)
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

	ctx := context.Background()

	tp, err := xrayconfig.NewTracerProvider(ctx)
	if err != nil {
		fmt.Printf("error creating tracer provider: %v", err)
	}

	defer func(ctx context.Context) {
		err := tp.Shutdown(ctx)
		if err != nil {
			fmt.Printf("error shutting down tracer provider: %v", err)
		}
	}(ctx)

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(xray.Propagator{})

	lambda.Start(otellambda.InstrumentHandler(Handler, xrayconfig.WithRecommendedOptions(tp)...))

}

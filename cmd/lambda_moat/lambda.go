package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/safetorun/PromptDefender/badwords"
	"github.com/safetorun/PromptDefender/badwords_embeddings"
	"github.com/safetorun/PromptDefender/embeddings"
	"github.com/safetorun/PromptDefender/internal/base_aws"
	"github.com/safetorun/PromptDefender/moat"
	"github.com/safetorun/PromptDefender/pii_aws"
	"go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-lambda-go/otellambda"
	"go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-lambda-go/otellambda/xrayconfig"
	"go.opentelemetry.io/contrib/propagators/aws/xray"
	"go.opentelemetry.io/otel"
	"os"
)

var (
	tracer = otel.Tracer("moat")
	meter  = otel.Meter("moat")
)

type MoatLambda struct {
	moatInstance *moat.Moat
}

func (m *MoatLambda) Handle(moatRequest MoatRequest) (*MoatResponse, error) {
	answer, err := m.moatInstance.CheckMoat(moat.PromptToCheck{
		Prompt:           moatRequest.Prompt,
		ScanPii:          moatRequest.ScanPii,
		XmlTagToCheckFor: moatRequest.XmlTag,
	},
	)

	containsPii := false
	if answer != nil && answer.PiiResult != nil {
		containsPii = answer.PiiResult.ContainsPii
	}

	if err != nil {
		return nil, err
	}

	var xmlEscaping *bool = nil

	if answer.XmlScannerResult != nil {
		xmlEscaping = &answer.XmlScannerResult.ContainsXmlEscaping
	}

	return &MoatResponse{ContainsPii: &containsPii, PotentialJailbreak: &answer.ContainsBadWords, PotentialXmlEscaping: xmlEscaping}, nil
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	ctx, span := tracer.Start(ctx, "moat_setup")

	openAIKey, exists := os.LookupEnv("open_ai_api_key")

	println("Received request for moat lambda")

	if !exists {
		return events.APIGatewayProxyResponse{StatusCode: 400}, fmt.Errorf("error with configuration")
	}

	addAllConfigurations := func(c *moat.Moat) error {
		c.PiiScanner = pii_aws.New()
		c.BadWordsCheck = badwords.New(badwords_embeddings.New(embeddings.New(openAIKey)))
		c.XmlEscapingScanner = moat.NewBasicXmlEscapingScaner()

		return nil
	}

	moatInstance, err := moat.New(addAllConfigurations)
	moatLambda := MoatLambda{
		moatInstance: moatInstance,
	}

	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400}, err
	}

	span.End()
	ctx, span = tracer.Start(ctx, "moat_handler_exec")
	defer span.End()

	response, err := base_aws.BaseHandler[MoatRequest, MoatResponse](request, &moatLambda)

	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400}, err
	}

	return response, nil
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

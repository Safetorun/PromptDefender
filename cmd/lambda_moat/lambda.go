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
	"github.com/safetorun/PromptDefender/tracer"
	"go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-lambda-go/otellambda"
	"go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-lambda-go/otellambda/xrayconfig"
	"go.opentelemetry.io/contrib/propagators/aws/xray"
	"go.opentelemetry.io/otel"
	"log"
	"os"
)

var (
	moat_tracer = otel.Tracer("moat")
	meter       = otel.Meter("moat")
)

type MoatLambda struct {
	moatInstance *moat.Moat
	context      context.Context
}

func (m *MoatLambda) Handle(moatRequest MoatRequest) (*MoatResponse, error) {

	t := tracer.NewTracer(m.context, moat_tracer)

	answer, err := m.moatInstance.CheckMoat(moat.PromptToCheck{
		Prompt:           moatRequest.Prompt,
		ScanPii:          moatRequest.ScanPii,
		XmlTagToCheckFor: moatRequest.XmlTag,
	},
		t,
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

func CheckSuspiciousUser(ctx context.Context, userId string, request events.APIGatewayProxyRequest) (*bool, error) {
	_, span := moat_tracer.Start(ctx, "check_suspicious_user")
	defer span.End()

	suspiciousUser := NewRemote(request.RequestContext.Identity.APIKey)
	isSuspicous, err := suspiciousUser.CheckSuspiciousUser(ctx, userId)

	if err != nil {
		return nil, err
	}

	return isSuspicous, nil
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	ctx, span := moat_tracer.Start(ctx, "moat_setup")

	openAIKey, exists := os.LookupEnv("open_ai_api_key")

	log.Default().Println("Received request for moat lambda")

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
		context:      ctx,
	}

	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400}, err
	}

	span.End()
	ctx, span = moat_tracer.Start(ctx, "moat_handler_exec")
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

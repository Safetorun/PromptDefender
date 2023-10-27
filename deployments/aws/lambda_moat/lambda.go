package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/safetorun/PromptDefender/aws/base_aws"
	"github.com/safetorun/PromptDefender/badwords"
	"github.com/safetorun/PromptDefender/badwords_embeddings"
	"github.com/safetorun/PromptDefender/embeddings"
	"github.com/safetorun/PromptDefender/moat"
	"github.com/safetorun/PromptDefender/pii_aws"
	"os"
)

type MoatLambda struct {
	moatInstance *moat.Moat
}

func (m *MoatLambda) Handle(moatRequest MoatRequest) (*MoatResponse, error) {
	answer, err := m.moatInstance.CheckMoat(moat.PromptToCheck{
		Prompt:  moatRequest.Prompt,
		ScanPii: moatRequest.ScanPii,
	},
	)

	containsPii := false
	if answer != nil && answer.PiiResult != nil {
		containsPii = answer.PiiResult.ContainsPii
	}

	if err != nil {
		return nil, err
	}

	return &MoatResponse{ContainsPii: &containsPii, PotentialJailbreak: &answer.ContainsBadWords}, nil
}

func Handler(_ context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	openAIKey, exists := os.LookupEnv("open_ai_api_key")

	if !exists {
		return events.APIGatewayProxyResponse{StatusCode: 400}, fmt.Errorf("error with configuration")
	}

	moatLambda := MoatLambda{
		moat.New(
			pii_aws.New(),
			badwords.New(badwords_embeddings.New(embeddings.New(openAIKey))),
		),
	}

	response, err := base_aws.BaseHandler[MoatRequest, MoatResponse](request, &moatLambda)

	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400}, err
	}

	return response, nil
}

func main() {
	lambda.Start(Handler)
}

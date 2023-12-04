package integration_test_harness

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func RequestToKeep(ctx context.Context) (context.Context, error) {
	return context.WithValue(ctx, RequestKey, &KeepRequest{}), nil
}

func SetRandomiseXmlTag(ctx context.Context, randomiseTag string) (context.Context, error) {

	randomise, err := strconv.ParseBool(randomiseTag)
	if err != nil {
		return nil, err
	}
	request := ctx.Value(RequestKey).(*KeepRequest)
	request.RandomiseXmlTag = &randomise
	return ctx, nil
}

func ValidateResponseXml(context context.Context, randomOrUserInput string) error {

	response := context.Value(ResponseKey).(*KeepResponse)

	if randomOrUserInput == "random" {
		if response.XmlTag == "user_input" {
			return errors.New("xml tag should not be random")
		}
	} else if response.XmlTag != randomOrUserInput {
		return errors.New("xml tag should be " + randomOrUserInput + " but was " + response.XmlTag)
	}

	return nil
}

func ShieldedPromptContains(context context.Context, containsText string) error {

	response := context.Value(ResponseKey).(*KeepResponse)

	if containsText == "random_tag" {
		containsText = "<" + response.XmlTag + ">"
	}

	if strings.Contains(response.ShieldedPrompt, containsText) == false {
		return errors.New("shielded prompt does not contain " + containsText + "Full prompt response " + response.ShieldedPrompt)
	}

	return nil
}

func SendRequestKeep(ctx context.Context) (context.Context, error) {
	gClient, err := CreateClient()

	if err != nil {
		return nil, err
	}

	if ctx.Value(RequestKey) == nil {
		return ctx, errors.New("request is nil")
	}

	request, ok := ctx.Value(RequestKey).(*KeepRequest)

	if ok == false {
		return ctx, errors.New("request is not castable to MoatRequest")
	}

	response, err := gClient.BuildKeepWithResponse(context.Background(), *request)

	if err != nil {
		return ctx, fmt.Errorf("got error (%s) when building shield", err)
	}

	if response.StatusCode() != 200 {
		return ctx, errors.New(fmt.Sprintf("error processing request %v", response))
	}

	if response.JSON200 == nil {
		return ctx, errors.New("response is nil")
	}

	return context.WithValue(ctx, ResponseKey, response.JSON200), nil

}

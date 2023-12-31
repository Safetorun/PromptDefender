package integration_test_harness

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

const RequestKey = "request"
const ResponseKey = "response"

func SendRequestToMoat(ctx context.Context) (context.Context, error) {
	gClient, err := CreateClient()

	if err != nil {
		return nil, err
	}

	if ctx.Value(RequestKey) == nil {
		return ctx, errors.New("request is nil")
	}

	request, ok := ctx.Value(RequestKey).(*MoatRequest)

	if ok == false {
		return ctx, errors.New("request is not castable to MoatRequest")
	}

	response, err := gClient.BuildShieldWithResponse(context.Background(), *request)

	if err != nil {
		return ctx, fmt.Errorf("got error (%s) when building shield", err)
	}

	if response.StatusCode() != 200 {
		return ctx, errors.New("error processing request")
	}

	if response.JSON200 == nil {
		return ctx, errors.New("response is nil")
	}

	return context.WithValue(ctx, ResponseKey, response.JSON200), nil

}

func RequestToMoat(ctx context.Context) (context.Context, error) {
	return context.WithValue(ctx, RequestKey, &MoatRequest{}), nil
}

func CreateClient() (*ClientWithResponses, error) {

	apiKey, exists := os.LookupEnv("DEFENDER_API_KEY")

	if exists == false {
		return nil, errors.New("DEFENDER_API_KEY not set")
	}

	url, exists := os.LookupEnv("URL")

	if exists == false {
		return nil, errors.New("URL not set")
	}

	addApiKey := func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, func(ctx context.Context, req *http.Request) error {
			req.Header.Add("x-api-key", apiKey)
			return nil
		})

		return nil
	}

	client, err := NewClientWithResponses(url, addApiKey)
	return client, err
}

func SetPiiDetection(ctx context.Context, enablePii string) (context.Context, error) {

	pii, err := strconv.ParseBool(enablePii)
	if err != nil {
		return nil, err
	}
	request := ctx.Value(RequestKey).(*MoatRequest)
	request.ScanPii = pii
	return ctx, nil
}

func SetPromptBody(ctx context.Context, prompt string) (context.Context, error) {
	request := ctx.Value(RequestKey).(*MoatRequest)
	request.Prompt = prompt
	return ctx, nil
}

func ValidateResponseXmlTagIsNil(context context.Context) error {

	response := context.Value(ResponseKey).(*MoatResponse)
	if response.PotentialXmlEscaping != nil {
		return errors.New("xml tag  should be nil")
	}

	return nil
}

func ValidateResponseXmlTag(context context.Context, xmlTag string) error {
	detected, err := strconv.ParseBool(xmlTag)
	if err != nil {
		return err
	}

	response := context.Value(ResponseKey).(*MoatResponse)
	if *response.PotentialXmlEscaping != detected {
		return errors.New("xml tag not set correctly")
	}

	return nil
}

func ValidateResponseDetectedPii(context context.Context, piiDetected string) error {

	detected, err := strconv.ParseBool(piiDetected)
	if err != nil {
		return err
	}

	response := context.Value(ResponseKey).(*MoatResponse)
	if *response.ContainsPii != detected {
		return errors.New("pii detected not set correctly")
	}

	return nil
}

func SetXmlTag(ctx context.Context, tag string) (context.Context, error) {
	request := ctx.Value(RequestKey).(*MoatRequest)
	request.XmlTag = &tag
	return ctx, nil
}

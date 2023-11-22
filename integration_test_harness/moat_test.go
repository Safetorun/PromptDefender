package integration_test_harness

import (
	"context"
	"errors"
	"github.com/cucumber/godog"
	"net/http"
	"os"
	"strconv"
	"testing"
)

const RequestKey = "request"
const ResponseKey = "response"

func sendRequest(ctx context.Context) (context.Context, error) {
	gClient, _ := createClient()
	response, err := gClient.BuildShieldWithResponse(ctx, *ctx.Value(RequestKey).(*MoatRequest))

	if err != nil {
		return nil, err
	}

	if response.StatusCode() != 200 {
		return nil, godog.ErrPending
	}

	return context.WithValue(ctx, ResponseKey, response.JSON200), nil

}

func requestToMoat(ctx context.Context) (context.Context, error) {
	return context.WithValue(ctx, RequestKey, &MoatRequest{}), nil
}

func createClient() (*ClientWithResponses, error) {
	addApiKey := func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, func(ctx context.Context, req *http.Request) error {
			req.Header.Add("x-api-key", os.Getenv("DEFENDER_API_KEY"))
			return nil
		})

		return nil
	}

	client, err := NewClientWithResponses("https://prompt.safetorun.com", addApiKey)
	return client, err
}

func setPiiDetection(ctx context.Context, enablePii string) (context.Context, error) {

	pii, err := strconv.ParseBool(enablePii)
	if err != nil {
		return nil, err
	}
	request := ctx.Value(RequestKey).(*MoatRequest)
	request.ScanPii = pii
	return ctx, nil
}

func setPromptBody(ctx context.Context, prompt string) (context.Context, error) {
	request := ctx.Value(RequestKey).(*MoatRequest)
	request.Prompt = prompt
	return ctx, nil
}

func validateResponseDetectedPii(context context.Context, piiDetected string) error {

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

func TestFeatures(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^I send a request to moat$`, requestToMoat)
	ctx.Step(`^I set PII detection to (true|false)$`, setPiiDetection)
	ctx.Step(`^the request is (.*)$`, setPromptBody)
	ctx.Step("^request is sent$", sendRequest)
	ctx.Step(`^Response should have PII detected set to (true|false)$`, validateResponseDetectedPii)
}

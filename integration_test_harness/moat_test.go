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

var (
	gClient *ClientWithResponses
)

func sendRequest(ctx context.Context) (context.Context, error) {
	response, err := gClient.BuildShieldWithResponse(ctx, *ctx.Value("request").(*MoatRequest))

	if err != nil {
		return nil, err
	}

	if response.StatusCode() != 200 {
		return nil, godog.ErrPending
	}

	return context.WithValue(ctx, "response", response.JSON200), nil

}

func requestToMoat(ctx context.Context) (context.Context, error) {
	addApiKey := func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, func(ctx context.Context, req *http.Request) error {
			req.Header.Add("x-api-key", os.Getenv("DEFENDER_API_KEY"))
			return nil
		})

		return nil
	}

	client, err := NewClientWithResponses("https://prompt.safetorun.com", addApiKey)
	gClient = client

	if err != nil {
		return nil, err
	}

	return context.WithValue(ctx, "request", &MoatRequest{}), nil
}

func setPiiDetection(ctx context.Context, enablePii string) (context.Context, error) {

	pii, err := strconv.ParseBool(enablePii)
	if err != nil {
		return nil, err
	}
	request := ctx.Value("request").(*MoatRequest)
	request.ScanPii = pii
	return ctx, nil
}

func setPromptBody(ctx context.Context, prompt string) (context.Context, error) {
	request := ctx.Value("request").(*MoatRequest)
	request.Prompt = prompt
	return ctx, nil
}

func validateResponseDetectedPii(context context.Context, piiDetected string) error {

	detected, err := strconv.ParseBool(piiDetected)
	if err != nil {
		return err
	}

	response := context.Value("response").(*MoatResponse)
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

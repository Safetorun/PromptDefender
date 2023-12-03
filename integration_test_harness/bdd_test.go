package integration_test_harness

import (
	"github.com/cucumber/godog"
	"testing"
)

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
	ctx.Step(`^I send a request to moat$`, RequestToMoat)
	ctx.Step(`^I send a request to keep$`, RequestToKeep)
	ctx.Step(`^I set PII detection to (true|false)$`, SetPiiDetection)
	ctx.Step(`^the request is (.*)$`, SetPromptBody)
	ctx.Step("^request is sent to moat$", SendRequestToMoat)
	ctx.Step(`^Response should have PII detected set to (true|false)$`, ValidateResponseDetectedPii)
	ctx.Step("^I set the XML tag to (.*)$", SetXmlTag)
	ctx.Step("^Response should detect XML tag escaping: (true|false)$", ValidateResponseXmlTag)
	ctx.Step("^Response should have XML tag escaping set to nil$", ValidateResponseXmlTagIsNil)
}

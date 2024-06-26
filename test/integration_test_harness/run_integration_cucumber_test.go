package integration_test_harness

import (
	"testing"

	"github.com/cucumber/godog"
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
	ctx.Step(`^I send a request to wall$`, RequestToWall)
	ctx.Step(`^I send a request to keep$`, RequestToKeep)
	ctx.Step(`^I set PII detection to (true|false)$`, SetPiiDetection)
	ctx.Step(`^the request is (.*)$`, SetPromptBody)
	ctx.Step("^request is sent to wall$", SendRequestToWall)
	ctx.Step(`^Response should have PII detected set to (true|false|nil)$`, ValidateResponseDetectedPii)
	ctx.Step("^I set the XML tag to (.*)$", SetXmlTag)
	ctx.Step("^Response should detect XML tag escaping: (true|false)$", ValidateResponseXmlTag)
	ctx.Step("^Response should have suspicious session input set to (true|false)$", ValidateSuspicousSessionInput)
	ctx.Step("^Suspicious user response should be (true|false)$", ValidateSuspicousUserInput)
	ctx.Step("^Response should have XML tag escaping set to nil$", ValidateResponseXmlTagIsNil)
	ctx.Step("^request is sent to keep", SendRequestKeep)
	ctx.Step("^I set randomise_xml_tag to (true|false$)", SetRandomiseXmlTag)
	ctx.Step("^Response should have xml_tag (user_input|random)$", ValidateResponseXml)
	ctx.Step("^Response should have shielded prompt containing (.*)$", ShieldedPromptContains)
	ctx.Step("^the keep request is (.*)$", SetKeepPrompt)
	ctx.Step("^the user_id is (.*)$", SetUserId)
	ctx.Step("^the session_id is (.*)$", SetSessionId)
	ctx.Step("^fast-check is (true|false)", SetFastCheck)
	ctx.Step("^check-badwords is (true|false)", SetCheckBadwords)
	ctx.Step("^Response should have shielded prompt containing (.*)$", ShieldedPromptContains)
	ctx.Step("^I create a suspicious user with the user_id (.*)$", CreateSuspiciousUser)
	ctx.Step("^I retrieve a list of all users$", RetrieveSuspiciousUsers)
	ctx.Step("^The list should contain the user_id (.*)$", ValidateUserIdContains)
	ctx.Step("^The list should not contain the user_id (.*)$", ValidateUserNotInList)
	ctx.Step("^I delete the user with the user_id (.*)$", DeleteSuspiciousUser)
	ctx.Step("^I retrieve the user with the user_id (.*)$", RetrieveSuspiciousUser)
	ctx.Step("^I should get a 404 error", Validate404Error)
	ctx.Step("^Response should have jailbreak detected set to (true|false)", ValidateResponsePotentialJailbreak)
}

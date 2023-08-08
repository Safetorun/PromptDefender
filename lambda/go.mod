module github.com/safetorun/PromptShield/lambda

go 1.20

replace github.com/safetorun/PromptShield/app => ../app

replace github.com/safetorun/PromptShield/aiprompt => ../aiprompt

require (
	github.com/aws/aws-lambda-go v1.41.0
	github.com/safetorun/PromptShield/aiprompt v0.0.0-00010101000000-000000000000
	github.com/safetorun/PromptShield/app v0.0.0-00010101000000-000000000000
)

module github.com/safetorun/PromptShield/deployments/aws/lambda_keep

go 1.20

replace github.com/safetorun/PromptShield/app => ../../../app

replace github.com/safetorun/PromptShield/aiprompt => ../../../aiprompt

replace github.com/safetorun/PromptShield/pii => ../../../pii

replace github.com/safetorun/PromptShield/pii_aws => ../../../pii_aws

replace github.com/safetorun/PromptShield/prompt => ../../../prompt

require (
	github.com/aws/aws-lambda-go v1.41.0
	github.com/safetorun/PromptShield/aiprompt v0.0.0-20231013064351-91b52b94a017
	github.com/safetorun/PromptShield/app v0.0.0-20231003080446-3884f5ee0afd
)

require github.com/safetorun/PromptShield/prompt v0.0.0-20231013064101-13fd645292ce // indirect

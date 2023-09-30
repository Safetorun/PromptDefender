module github.com/safetorun/PromptShield/lambda

go 1.20

replace github.com/safetorun/PromptShield/app => ../app

replace github.com/safetorun/PromptShield/aiprompt => ../aiprompt

replace github.com/safetorun/PromptShield/pii => ../pii

replace github.com/safetorun/PromptShield/pii_aws => ../pii_aws

require (
	github.com/aws/aws-lambda-go v1.41.0
	github.com/safetorun/PromptShield/aiprompt v0.0.0-20230930091917-0b3812539293
	github.com/safetorun/PromptShield/app v0.0.0-00010101000000-000000000000
	github.com/safetorun/PromptShield/pii_aws v0.0.0-00010101000000-000000000000
)

require (
	github.com/aws/aws-sdk-go v1.45.19 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/safetorun/PromptShield/pii v0.0.0-20230930091917-0b3812539293 // indirect
)

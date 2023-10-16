module github.com/safetorun/PromptShield/deployments/aws/lambda_moat

go 1.20

replace github.com/safetorun/PromptShield/app => ../../../app

replace github.com/safetorun/PromptShield/aiprompt => ../../../aiprompt

replace github.com/safetorun/PromptShield/pii => ../../../pii

replace github.com/safetorun/PromptShield/pii_aws => ../../../pii_aws

replace github.com/safetorun/PromptShield/prompt => ../../../prompt

require (
	github.com/aws/aws-lambda-go v1.41.0
	github.com/safetorun/PromptShield/aiprompt v0.0.0-20231014075714-cbf445fc8e67
	github.com/safetorun/PromptShield/app v0.0.0-20231014075714-cbf445fc8e67
	github.com/safetorun/PromptShield/pii_aws v0.0.0-20231014075714-cbf445fc8e67
)

require (
	github.com/aws/aws-sdk-go v1.45.25 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/safetorun/PromptShield/pii v0.0.0-20231014075714-cbf445fc8e67 // indirect
	github.com/safetorun/PromptShield/prompt v0.0.0-20231014075714-cbf445fc8e67 // indirect
)

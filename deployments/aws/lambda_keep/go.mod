module github.com/safetorun/PromptDefender/deployments/aws/lambda_keep

go 1.20

replace github.com/safetorun/PromptDefender/app => ../../../app

replace github.com/safetorun/PromptDefender/aiprompt => ../../../aiprompt

replace github.com/safetorun/PromptDefender/pii => ../../../pii

replace github.com/safetorun/PromptDefender/pii_aws => ../../../pii_aws

replace github.com/safetorun/PromptDefender/prompt => ../../../prompt

replace github.com/safetorun/PromptDefender/keep => ../../../keep

require (
	github.com/aws/aws-lambda-go v1.41.0
	github.com/safetorun/PromptDefender/aiprompt v0.0.0-20231014075714-cbf445fc8e67
	github.com/safetorun/PromptDefender/keep v0.0.0-00010101000000-000000000000
)

require (
	github.com/safetorun/PromptDefender/prompt v0.0.0-20231014075714-cbf445fc8e67 // indirect
	github.com/stretchr/testify v1.8.4 // indirect
)

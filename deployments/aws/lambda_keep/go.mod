module github.com/safetorun/PromptDefender/deployments/aws/lambda_keep

go 1.20

replace github.com/safetorun/PromptDefender/aws/base_aws => ../base_aws

replace github.com/safetorun/PromptDefender/aiprompt => ../../../aiprompt

replace github.com/safetorun/PromptDefender/prompt => ../../../prompt

replace github.com/safetorun/PromptDefender/keep => ../../../keep

require (
	github.com/aws/aws-lambda-go v1.41.0
	github.com/aws/aws-sdk-go-v2 v1.23.0
	github.com/aws/aws-sdk-go-v2/config v1.25.1
	github.com/aws/aws-sdk-go-v2/service/sqs v1.28.1
	github.com/safetorun/PromptDefender/aiprompt v0.0.0-20231017173944-0a5da2b1ee56
	github.com/safetorun/PromptDefender/aws/base_aws v0.0.0-00010101000000-000000000000
	github.com/safetorun/PromptDefender/keep v0.0.0-20231017173944-0a5da2b1ee56
)

require (
	github.com/aws/aws-sdk-go-v2/credentials v1.16.1 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.14.4 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.2.3 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.5.3 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.7.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.10.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.10.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.17.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.19.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.25.2 // indirect
	github.com/aws/smithy-go v1.17.0 // indirect
	github.com/safetorun/PromptDefender/prompt v0.0.0-20231017173944-0a5da2b1ee56 // indirect
)

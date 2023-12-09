module github.com/safetorun/PromptDefender/deployments/aws/lambda_keep

go 1.20

replace github.com/safetorun/PromptDefender/aws/base_aws => ./../../internal/base_aws

replace github.com/safetorun/PromptDefender/aiprompt => ./../../aiprompt

replace github.com/safetorun/PromptDefender/keep => ./../../keep

require (
	github.com/aws/aws-lambda-go v1.41.0
	github.com/aws/aws-sdk-go-v2 v1.23.5
	github.com/aws/aws-sdk-go-v2/config v1.25.11
	github.com/aws/aws-sdk-go-v2/service/sqs v1.29.2
	github.com/safetorun/PromptDefender/aiprompt v0.0.0-20231204191015-931fb010b80b
	github.com/safetorun/PromptDefender/aws/base_aws v0.0.0-00010101000000-000000000000
	github.com/safetorun/PromptDefender/keep v0.0.0-20231204191015-931fb010b80b
)

require (
	github.com/aws/aws-sdk-go-v2/credentials v1.16.9 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.14.9 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.2.8 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.5.8 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.7.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.10.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.10.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.18.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.21.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.26.2 // indirect
	github.com/aws/smithy-go v1.18.1 // indirect
)

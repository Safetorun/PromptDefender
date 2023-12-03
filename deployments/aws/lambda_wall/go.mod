module github.com/safetorun/PromptDefender/deployments/aws/lambda_wall

go 1.20

replace github.com/safetorun/PromptDefender/aws/base_aws => ../base_aws

replace github.com/safetorun/PromptDefender/aiprompt => ../../../aiprompt

replace github.com/safetorun/PromptDefender/wall => ../../../wall

require (
	github.com/aws/aws-lambda-go v1.41.0
	github.com/safetorun/PromptDefender/aiprompt v0.0.0-20231203123834-51c0ba01645f
	github.com/safetorun/PromptDefender/aws/base_aws v0.0.0-00010101000000-000000000000
	github.com/safetorun/PromptDefender/wall v0.0.0-20231203123834-51c0ba01645f
)

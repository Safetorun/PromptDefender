module github.com/safetorun/PromptDefender/deployments/aws/lambda_moat

go 1.20

replace github.com/safetorun/PromptDefender/moat => ../../../moat

replace github.com/safetorun/PromptDefender/aiprompt => ../../../aiprompt

replace github.com/safetorun/PromptDefender/pii => ../../../pii

replace github.com/safetorun/PromptDefender/pii_aws => ../../../pii_aws

require (
	github.com/aws/aws-lambda-go v1.41.0
	github.com/safetorun/PromptDefender/aiprompt v0.0.0-20231017075420-f34296ffedcf
	github.com/safetorun/PromptDefender/moat v0.0.0-20231017114355-94a066de9d10
	github.com/safetorun/PromptDefender/pii_aws v0.0.0-20231017075420-f34296ffedcf
)

require (
	github.com/aws/aws-sdk-go v1.45.26 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/safetorun/PromptDefender/pii v0.0.0-20231017075420-f34296ffedcf // indirect
)

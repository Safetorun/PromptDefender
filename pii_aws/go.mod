module github.com/safetorun/PromptDefender/pii_aws

go 1.20

replace github.com/safetorun/PromptDefender/pii => ../pii

require (
	github.com/aws/aws-sdk-go v1.48.11
	github.com/safetorun/PromptDefender/pii v0.0.0-20231203123834-51c0ba01645f
)

require github.com/jmespath/go-jmespath v0.4.0 // indirect

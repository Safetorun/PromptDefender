module github.com/safetorun/PromptDefender/pii_aws

go 1.20

replace github.com/safetorun/PromptDefender/pii => ../pii

require (
	github.com/aws/aws-sdk-go v1.45.26
	github.com/safetorun/PromptDefender/pii v0.0.0-20231017075420-f34296ffedcf
)

require github.com/jmespath/go-jmespath v0.4.0 // indirect

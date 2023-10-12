module github.com/safetorun/PromptShield/pii_aws

go 1.20

replace github.com/safetorun/PromptShield/pii => ../pii

require (
	github.com/aws/aws-sdk-go v1.45.24
	github.com/safetorun/PromptShield/pii v0.0.0-20231003080446-3884f5ee0afd
)

require github.com/jmespath/go-jmespath v0.4.0 // indirect

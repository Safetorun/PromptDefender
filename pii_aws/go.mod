module github.com/safetorun/PromptShield/pii_aws

go 1.20

replace github.com/safetorun/PromptShield/pii => ../pii

require (
	github.com/aws/aws-sdk-go v1.45.1
	github.com/safetorun/PromptShield/pii v0.0.0-00010101000000-000000000000
)

require github.com/jmespath/go-jmespath v0.4.0 // indirect

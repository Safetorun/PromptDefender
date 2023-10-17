module github.com/safetorun/PromptDefender/pii_aws

go 1.20

replace github.com/safetorun/PromptDefender/pii => ../pii

require (
	github.com/aws/aws-sdk-go v1.45.27
	github.com/safetorun/PromptDefender/pii v0.0.0-20231017173944-0a5da2b1ee56
)

require github.com/jmespath/go-jmespath v0.4.0 // indirect

module github.com/safetorun/PromptDefender/pii_aws

go 1.20

replace github.com/safetorun/PromptDefender/pii => ../pii

require (
	github.com/aws/aws-sdk-go v1.45.25
	github.com/safetorun/PromptDefender/pii v0.0.0-20231014075714-cbf445fc8e67
)

require github.com/jmespath/go-jmespath v0.4.0 // indirect

module github.com/safetorun/PromptDefender/deployments/aws/lambda_keep

go 1.20

replace github.com/safetorun/PromptDefender/aiprompt => ../../../aiprompt

replace github.com/safetorun/PromptDefender/pii => ../../../pii

replace github.com/safetorun/PromptDefender/pii_aws => ../../../pii_aws

replace github.com/safetorun/PromptDefender/prompt => ../../../prompt

replace github.com/safetorun/PromptDefender/keep => ../../../keep

require (
	github.com/aws/aws-lambda-go v1.41.0
	github.com/safetorun/PromptDefender/aiprompt v0.0.0-20231017173944-0a5da2b1ee56
	github.com/safetorun/PromptDefender/keep v0.0.0-20231017173944-0a5da2b1ee56
)

require github.com/safetorun/PromptDefender/prompt v0.0.0-20231017173944-0a5da2b1ee56 // indirect

module github.com/safetorun/PromptDefender/deployments/aws/lambda_wall

go 1.20

replace github.com/safetorun/PromptDefender/aiprompt => ../../../aiprompt

replace github.com/safetorun/PromptDefender/prompt => ../../../prompt

replace github.com/safetorun/PromptDefender/wall => ../../../wall

require (
	github.com/aws/aws-lambda-go v1.41.0
	github.com/safetorun/PromptDefender/aiprompt v0.0.0-20231017173944-0a5da2b1ee56
	github.com/safetorun/PromptDefender/wall v0.0.0-20231017173944-0a5da2b1ee56
)
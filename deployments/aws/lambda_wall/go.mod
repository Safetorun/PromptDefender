module github.com/safetorun/PromptDefender/deployments/aws/lambda_wall

go 1.20

replace github.com/safetorun/PromptDefender/aiprompt => ../../../aiprompt

replace github.com/safetorun/PromptDefender/prompt => ../../../prompt

replace github.com/safetorun/PromptDefender/wall => ../../../wall

require (
	github.com/aws/aws-lambda-go v1.41.0
	github.com/safetorun/PromptDefender/aiprompt v0.0.0-20231017114830-c0695821fa78
	github.com/safetorun/PromptDefender/wall v0.0.0-00010101000000-000000000000
)

module github.com/safetorun/PromptDefender/deployments/aws/lambda_moat

go 1.20

replace github.com/safetorun/PromptDefender/moat => ./../../moat

replace github.com/safetorun/PromptDefender/pii => ./../../pii

replace github.com/safetorun/PromptDefender/badwords => ./../../badwords

replace github.com/safetorun/PromptDefender/badwords_embeddings => ./../../badwords_embeddings

replace github.com/safetorun/PromptDefender/embeddings => ./../../embeddings

replace github.com/safetorun/PromptDefender/pii_aws => ./../../pii_aws

require (
	github.com/aws/aws-lambda-go v1.41.0
	github.com/safetorun/PromptDefender/badwords v0.0.0-20231204191015-931fb010b80b
	github.com/safetorun/PromptDefender/badwords_embeddings v0.0.0-20231204191015-931fb010b80b
	github.com/safetorun/PromptDefender/embeddings v0.0.0-20231204191015-931fb010b80b
	github.com/safetorun/PromptDefender/moat v0.0.0-20231204191015-931fb010b80b
	github.com/safetorun/PromptDefender/pii_aws v0.0.0-20231204191015-931fb010b80b
)

require (
	github.com/aws/aws-sdk-go v1.48.13 // indirect
	github.com/drewlanenga/govector v0.0.0-20220726163947-b958ac08bc93 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/rogpeppe/go-internal v1.11.0 // indirect
	github.com/safetorun/PromptDefender/pii v0.0.0-20231204191015-931fb010b80b // indirect
	github.com/sashabaranov/go-openai v1.17.9 // indirect
)

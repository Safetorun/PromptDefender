module github.com/safetorun/PromptDefender/deployments/aws/lambda_moat

go 1.20

require (
	github.com/aws/aws-lambda-go v1.43.0
	github.com/safetorun/PromptDefender/badwords v0.0.0-20231210112259-15b98f65cf67
	github.com/safetorun/PromptDefender/badwords_embeddings v0.0.0-20231210112259-15b98f65cf67
	github.com/safetorun/PromptDefender/embeddings v0.0.0-20231210112259-15b98f65cf67
	github.com/safetorun/PromptDefender/internal/base_aws v0.0.0-20231210114542-bda8793eb24f
	github.com/safetorun/PromptDefender/moat v0.0.0-20231210112259-15b98f65cf67
	github.com/safetorun/PromptDefender/pii_aws v0.0.0-20231210112259-15b98f65cf67
)

require (
	github.com/aws/aws-sdk-go v1.49.16 // indirect
	github.com/drewlanenga/govector v0.0.0-20220726163947-b958ac08bc93 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/rogpeppe/go-internal v1.11.0 // indirect
	github.com/safetorun/PromptDefender/pii v0.0.0-20231210081706-40db07878111 // indirect
	github.com/sashabaranov/go-openai v1.17.10 // indirect
)

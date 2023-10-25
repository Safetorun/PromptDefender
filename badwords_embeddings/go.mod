module github.com/safetorun/PromptDefender/badwords_embeddings

replace github.com/safetorun/PromptDefender/badwords => ../badwords

replace github.com/safetorun/PromptDefender/embeddings => ../embeddings

go 1.20

require (
	github.com/drewlanenga/govector v0.0.0-20220726163947-b958ac08bc93
	github.com/safetorun/PromptDefender/badwords v0.0.0-00010101000000-000000000000
	github.com/safetorun/PromptDefender/embeddings v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.8.4
)

require (
	github.com/bmizerany/assert v0.0.0-20160611221934-b7ed37b82869 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/sashabaranov/go-openai v1.16.0 // indirect
	github.com/stretchr/objx v0.5.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

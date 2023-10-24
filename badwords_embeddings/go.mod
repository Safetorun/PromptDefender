module github.com/safetorun/PromptDefender/badwords_embeddings

replace github.com/safetorun/PromptDefender/badwords => ../badwords

replace github.com/safetorun/PromptDefender/embeddings => ../embeddings

go 1.20

require (
	github.com/safetorun/PromptDefender/badwords v0.0.0-00010101000000-000000000000
	github.com/safetorun/PromptDefender/embeddings v0.0.0-00010101000000-000000000000
)

require github.com/sashabaranov/go-openai v1.16.0 // indirect

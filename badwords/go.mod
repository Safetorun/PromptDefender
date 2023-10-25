module github.com/safetorun/PromptDefender/badwords

replace (
	github.com/safetorun/PromptDefender/badwords => ../badwords
	github.com/safetorun/PromptDefender/embeddings => ../embeddings
)

go 1.20

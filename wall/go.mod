module github.com/safetorun/PromptDefender/wall

replace github.com/safetorun/PromptDefender/aiprompt => ../aiprompt

replace github.com/safetorun/PromptDefender/pii => ../pii

replace github.com/safetorun/PromptDefender/prompt => ../prompt

go 1.20

require (
	github.com/safetorun/PromptDefender/aiprompt v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.8.4
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

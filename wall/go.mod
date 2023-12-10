module github.com/safetorun/PromptDefender/wall

replace github.com/safetorun/PromptDefender/aiprompt => ./../internal/aiprompt

replace github.com/safetorun/PromptDefender/pii => ./../internal/pii

go 1.20

require (
	github.com/safetorun/PromptDefender/aiprompt v0.0.0-20231204191015-931fb010b80b
	github.com/stretchr/testify v1.8.4
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

module github.com/safetorun/PromptDefender/keep

replace github.com/safetorun/PromptDefender/aiprompt => ../aiprompt

replace github.com/safetorun/PromptDefender/prompt => ../prompt

require (
	github.com/safetorun/PromptDefender/aiprompt v0.0.0-20231014075714-cbf445fc8e67
	github.com/safetorun/PromptDefender/prompt v0.0.0-20231014075714-cbf445fc8e67
	github.com/stretchr/testify v1.8.4
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

go 1.20

module github.com/safetorun/PromptDefender/keep

replace github.com/safetorun/PromptDefender/aiprompt => ../aiprompt

require (
	github.com/safetorun/PromptDefender/aiprompt v0.0.0-20231203123834-51c0ba01645f
	github.com/stretchr/testify v1.8.4
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

go 1.20

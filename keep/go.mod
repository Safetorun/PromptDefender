module github.com/safetorun/PromptDefender/keep

replace github.com/safetorun/PromptDefender/aiprompt => ../aiprompt

replace github.com/safetorun/PromptDefender/prompt => ../prompt

require (
	github.com/safetorun/PromptDefender/aiprompt v0.0.0-20231017075420-f34296ffedcf
	github.com/safetorun/PromptDefender/prompt v0.0.0-20231017114355-94a066de9d10
	github.com/stretchr/testify v1.8.4
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

go 1.20

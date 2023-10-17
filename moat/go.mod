module github.com/safetorun/PromptDefender/moat

replace github.com/safetorun/PromptDefender/pii => ../pii

replace github.com/safetorun/PromptDefender/prompt => ../prompt

go 1.20

require (
	github.com/safetorun/PromptDefender/pii v0.0.0-20231017114830-c0695821fa78
	github.com/stretchr/testify v1.8.4
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

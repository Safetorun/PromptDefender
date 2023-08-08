module github.com/safetorun/PromptShield/app

go 1.17

replace github.com/safetorun/PromptShield/aiprompt => ../aiprompt

require github.com/safetorun/PromptShield/aiprompt v0.0.0-00010101000000-000000000000

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/objx v0.5.0 // indirect
	github.com/stretchr/testify v1.8.2
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

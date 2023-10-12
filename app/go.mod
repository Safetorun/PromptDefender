module github.com/safetorun/PromptShield/app

go 1.17

replace github.com/safetorun/PromptShield/aiprompt => ../aiprompt

replace github.com/safetorun/PromptShield/pii => ../pii

replace github.com/safetorun/PromptShield/pii_aws => ../pii_aws

replace github.com/safetorun/PromptShield/prompt => ../prompt

require (
	github.com/safetorun/PromptShield/aiprompt v0.0.0-20231003080446-3884f5ee0afd
	github.com/safetorun/PromptShield/prompt v0.0.0-20231003080446-3884f5ee0afd
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/stretchr/objx v0.5.1 // indirect
	github.com/stretchr/testify v1.8.4
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

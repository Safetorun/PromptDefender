module github.com/safetorun/PromptDefender/integration_test_harness

replace github.com/safetorun/PromptDefender/aiprompt => ../aiprompt

go 1.20

require (
	github.com/cucumber/godog v0.13.0
	github.com/safetorun/PromptDefender/aiprompt v0.0.0-20231203123834-51c0ba01645f
	github.com/safetorun/PromptDefender/keep v0.0.0-20231203123834-51c0ba01645f
	github.com/stretchr/testify v1.8.4
)

require (
	github.com/cucumber/gherkin/go/v26 v26.2.0 // indirect
	github.com/cucumber/messages/go/v21 v21.0.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/gofrs/uuid v4.4.0+incompatible // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-memdb v1.3.4 // indirect
	github.com/hashicorp/golang-lru v1.0.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/safetorun/PromptDefender/prompt v0.0.0-20231203123834-51c0ba01645f // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

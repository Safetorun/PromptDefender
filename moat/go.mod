module github.com/safetorun/PromptDefender/moat

replace github.com/safetorun/PromptDefender/aiprompt => ../aiprompt

replace github.com/safetorun/PromptDefender/pii => ../pii

replace github.com/safetorun/PromptDefender/prompt => ../prompt

go 1.20

require (
	github.com/safetorun/PromptDefender/aiprompt v0.0.0-00010101000000-000000000000
	github.com/safetorun/PromptDefender/pii v0.0.0-00010101000000-000000000000
)

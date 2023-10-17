module github.com/safetorun/PromptDefender/keep


replace github.com/safetorun/PromptDefender/aiprompt => ../aiprompt

replace github.com/safetorun/PromptDefender/prompt => ../prompt

require (
	github.com/safetorun/PromptDefender/aiprompt v0.0.0-20231014075714-cbf445fc8e67
	github.com/safetorun/PromptDefender/prompt v0.0.0-20231014075714-cbf445fc8e67
)


go 1.20

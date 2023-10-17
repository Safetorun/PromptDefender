package keep

import (
	"github.com/safetorun/PromptDefender/aiprompt"
	"github.com/safetorun/PromptDefender/prompt"
)

type Keep struct {
	openAi aiprompt.RemoteAIChecker
}

type StartingPrompt struct {
	Prompt string
}

type NewPrompt struct {
	NewPrompt string
}

func New(aiPrompt aiprompt.RemoteAIChecker) *Keep {
	return &Keep{
		openAi: aiPrompt,
	}
}

func (k *Keep) BuildKeep(startingPrompt StartingPrompt) (*NewPrompt, error) {

	builtPrompt := prompt.SmartPrompt(prompt.SmartPromptRequest{BasePrompt: startingPrompt.Prompt})

	response, err := k.openAi.CheckAI(builtPrompt)

	if err != nil {
		return nil, err
	}

	return &NewPrompt{NewPrompt: *response}, nil
}

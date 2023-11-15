package keep

import (
	"github.com/safetorun/PromptDefender/aiprompt"
	"github.com/safetorun/PromptDefender/prompt"
)

type KeepCallback *func(prompt string, newPrompt string)

type KeepOption func(*Keep)

type Keep struct {
	openAi aiprompt.RemoteAIChecker
}

type StartingPrompt struct {
	Prompt string
}

type NewPrompt struct {
	NewPrompt string
	Callback  KeepCallback
}

func New(aiPrompt aiprompt.RemoteAIChecker, options ...KeepOption) *Keep {
	k := &Keep{
		openAi: aiPrompt,
	}

	for _, opt := range options {
		opt(k)
	}

	return k
}

func (k *Keep) BuildKeep(startingPrompt StartingPrompt) (*NewPrompt, error) {

	builtPrompt := prompt.SmartPrompt(prompt.SmartPromptRequest{BasePrompt: startingPrompt.Prompt})

	response, err := k.openAi.CheckAI(builtPrompt)

	if err != nil {
		return nil, err
	}

	return &NewPrompt{NewPrompt: *response}, nil
}

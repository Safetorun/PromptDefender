package keep

import (
	"github.com/safetorun/PromptDefender/aiprompt"
	"github.com/safetorun/PromptDefender/prompt"
)

type Callback func(prompt string, newPrompt string) error

type KeepOption func(*Keep)

type Keep struct {
	openAi   aiprompt.RemoteAIChecker
	Callback *Callback
}

type StartingPrompt struct {
	Prompt string
}

type NewPrompt struct {
	NewPrompt string
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

	if k.Callback != nil {
		err := (*k.Callback)(startingPrompt.Prompt, *response)
		if err != nil {
			return nil, err
		}
	}

	return &NewPrompt{NewPrompt: *response}, nil
}

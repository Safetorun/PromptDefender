package wall

import (
	"github.com/safetorun/PromptDefender/aiprompt"
	"strconv"
)

type Wall struct {
	remotePrompter aiprompt.RemoteAIChecker
}

type PromptToCheck struct {
	Prompt string
}

type InjectionScore struct {
	Score float32
}

func New(prompter aiprompt.RemoteAIChecker) Wall {
	return Wall{
		remotePrompter: prompter,
	}
}

func (w *Wall) CheckWall(check PromptToCheck) (*InjectionScore, error) {
	return w.retrieveInjectionScore(check.Prompt)
}

// Send the prompt to AI in order to get a score for how likely it is to be an injection attack
func (w *Wall) retrieveInjectionScore(prompt string) (*InjectionScore, error) {
	promptForPiDetection := aiprompt.RenderPromptForPiDetection(prompt)

	answer, err := w.remotePrompter.CheckAI(promptForPiDetection)
	if err != nil {
		return nil, err
	}

	f, err := strconv.ParseFloat(*answer, 32)

	if err != nil {
		return nil, err
	}

	return &InjectionScore{float32(f)}, nil
}

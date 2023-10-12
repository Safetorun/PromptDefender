package app

import (
	"github.com/safetorun/PromptShield/aiprompt"
	"github.com/safetorun/PromptShield/prompt"
	"strconv"
)

type IApp interface {
	CheckAI(basePrompt string) (float32, error)
	BuildPromptDefense(basePrompt string) string
}

type App struct {
	openAi aiprompt.RemoteAIChecker
}

func New(openAi aiprompt.RemoteAIChecker) *App {
	return &App{openAi: openAi}
}

func (app *App) BuildPromptDefense(basePrompt string) (*string, error) {
	answer, err := app.openAi.CheckAI(prompt.SmartPrompt(prompt.SmartPromptRequest{BasePrompt: basePrompt}))
	if err != nil {
		return nil, err
	}

	return answer, nil

}

func (app *App) CheckAI(basePrompt string) (float32, error) {
	promptForPiDetection := aiprompt.RenderPromptForPiDetection(basePrompt)

	answer, err := app.openAi.CheckAI(promptForPiDetection)
	if err != nil {
		return -1, err
	}

	f, err := strconv.ParseFloat(*answer, 32)

	if err != nil {
		return -1, err
	}

	return float32(f), nil
}

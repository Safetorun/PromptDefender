package app

import (
	"strconv"

	"github.com/safetorun/PromptShield/aiprompt"
)

type IApp interface {
	CheckAI(basePrompt string) (float32, error)
}

type App struct {
	openAi aiprompt.RemoteAIChecker
}

func New(openAi aiprompt.RemoteAIChecker) *App {
	return &App{openAi: openAi}
}

func (app *App) CheckAI(basePrompt string) (float32, error) {
	prompt := aiprompt.RenderPromptForPiDetection(basePrompt)

	answer, err := app.openAi.CheckAI(prompt)
	if err != nil {
		return -1, err
	}

	f, err := strconv.ParseFloat(*answer, 32)

	if err != nil {
		return -1, err
	}

	return float32(f), nil
}

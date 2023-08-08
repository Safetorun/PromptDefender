package app

import (
	"github.com/safetorun/PromptShield/aiprompt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type OpenAIMock struct {
	mock.Mock
}

func (m *OpenAIMock) CheckAI(prompt string) (*string, error) {
	args := m.Called(prompt)
	return args.Get(0).(*string), args.Error(1)
}

func TestCheckAI(t *testing.T) {
	openAiMock := new(OpenAIMock)
	app := &App{openAi: openAiMock}

	basePrompt := "test prompt"
	expectedPrompt := aiprompt.RenderPromptForPiDetection(basePrompt)
	returnValue := "1.23"
	openAiMock.On("CheckAI", expectedPrompt).Return(&returnValue, nil)

	result, err := app.CheckAI(basePrompt)

	assert.Nil(t, err)
	assert.Equal(t, float32(1.23), result)

	openAiMock.AssertExpectations(t)
}

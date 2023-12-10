package wall

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

type MockAIChecker struct {
	checkAI func(string) (*string, error)
}

func (m *MockAIChecker) CheckAI(prompt string) (*string, error) {
	return m.checkAI(prompt)
}

func TestCheckMoat(t *testing.T) {
	mockAIChecker := &MockAIChecker{
		checkAI: func(_ string) (*string, error) {
			score := "0.9"
			return &score, nil
		},
	}

	mockAICheckerError := &MockAIChecker{
		checkAI: func(_ string) (*string, error) {
			return nil, errors.New("AI error")
		},
	}

	mockParseFloatError := &MockAIChecker{
		checkAI: func(_ string) (*string, error) {
			invalidFloat := "invalid"
			return &invalidFloat, nil
		},
	}

	t.Run("CheckMoat_Success", func(t *testing.T) {
		w := New(mockAIChecker)
		score, err := w.CheckWall(PromptToCheck{"test prompt"})
		assert.NoError(t, err)
		assert.Equal(t, float32(0.9), score.Score)
	})

	t.Run("CheckMoat_Error_AI", func(t *testing.T) {
		w := New(mockAICheckerError)
		_, err := w.CheckWall(PromptToCheck{"test prompt"})
		assert.Error(t, err)
	})

	t.Run("CheckMoat_Error_ParseFloat", func(t *testing.T) {
		w := New(mockParseFloatError)
		_, err := w.CheckWall(PromptToCheck{"test prompt"})
		assert.Error(t, err)
	})
}

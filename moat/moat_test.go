package moat

import (
	"errors"
	"github.com/safetorun/PromptDefender/pii"
	"github.com/stretchr/testify/assert"
	"testing"
)

type MockAIChecker struct{}

func (m *MockAIChecker) CheckAI(_ string) (*string, error) {
	score := "0.9"
	return &score, nil
}

type MockPIIScanner struct{}

func (m *MockPIIScanner) Scan(string) (*pii.ScanResult, error) {
	return &pii.ScanResult{ContainingPii: false}, nil
}

type MockPIIScannerError struct{}

func (m *MockPIIScannerError) Scan(_ string) (*pii.ScanResult, error) {
	return nil, errors.New("PII scan error")
}

func TestMoat(t *testing.T) {
	mockPIIScanner := &MockPIIScanner{}
	mockPIIScannerError := &MockPIIScannerError{}
	m := New(mockPIIScanner)

	t.Run("CheckMoat_WithPIIScan", func(t *testing.T) {
		promptToCheck := PromptToCheck{
			Prompt:  "test prompt",
			ScanPii: true,
		}

		result, err := m.CheckMoat(promptToCheck)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.NotNil(t, result.PiiResult)
	})

	t.Run("CheckMoat_WithoutPIIScan", func(t *testing.T) {
		promptToCheck := PromptToCheck{
			Prompt:  "test prompt",
			ScanPii: false,
		}

		result, err := m.CheckMoat(promptToCheck)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Nil(t, result.PiiResult)
	})

	t.Run("CheckMoat_PiiScannerError", func(t *testing.T) {
		m := New(mockPIIScannerError)
		promptToCheck := PromptToCheck{
			Prompt:  "test prompt",
			ScanPii: true,
		}

		_, err := m.CheckMoat(promptToCheck)
		assert.Error(t, err)
	})
}

package moat

import (
	"fmt"
	"github.com/safetorun/PromptDefender/aiprompt"
	"github.com/safetorun/PromptDefender/pii"
	"strconv"
)

type Moat struct {
	openAi     aiprompt.RemoteAIChecker
	piiScanner pii.Scanner
}

type PromptToCheck struct {
	Prompt  string
	ScanPii bool
}

type InjectionScore struct {
	Score float32
}

type PiiDetectionResult struct {
	ContainsPii bool
}

type CheckResult struct {
	InjectionScore *InjectionScore
	PiiResult      *PiiDetectionResult
}

func New(ai aiprompt.RemoteAIChecker, piiScanner pii.Scanner) *Moat {
	return &Moat{openAi: ai, piiScanner: piiScanner}
}

func (m *Moat) CheckMoat(check PromptToCheck) (*CheckResult, error) {
	injectionScore, err := m.retrieveInjectionScore(check.Prompt)

	if err != nil {
		return nil, err
	}

	var piiResult *PiiDetectionResult = nil

	if check.ScanPii {
		piiResult, err = m.retrievePiiScore(check.Prompt)

		if err != nil {
			return nil, err
		}
	}

	return &CheckResult{injectionScore, piiResult}, nil
}

func (m *Moat) retrievePiiScore(basePrompt string) (*PiiDetectionResult, error) {
	piiResult, err := m.piiScanner.Scan(basePrompt)

	if err != nil {
		fmt.Printf("Error scanning for PII %v\n", err)
		return nil, err
	}

	return &PiiDetectionResult{ContainsPii: piiResult.ContainingPii}, nil
}

// Send the prompt to AI in order to get a score for how likely it is to be an injection attack
func (m *Moat) retrieveInjectionScore(prompt string) (*InjectionScore, error) {
	promptForPiDetection := aiprompt.RenderPromptForPiDetection(prompt)

	answer, err := m.openAi.CheckAI(promptForPiDetection)
	if err != nil {
		return nil, err
	}

	f, err := strconv.ParseFloat(*answer, 32)

	if err != nil {
		return nil, err
	}

	return &InjectionScore{float32(f)}, nil
}

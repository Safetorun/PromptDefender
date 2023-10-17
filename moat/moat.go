package moat

import (
	"fmt"
	"github.com/safetorun/PromptDefender/pii"
)

type Moat struct {
	piiScanner pii.Scanner
}

type PromptToCheck struct {
	Prompt  string
	ScanPii bool
}

type PiiDetectionResult struct {
	ContainsPii bool
}

type CheckResult struct {
	PiiResult *PiiDetectionResult
}

func New(piiScanner pii.Scanner) *Moat {
	return &Moat{piiScanner: piiScanner}
}

func (m *Moat) CheckMoat(check PromptToCheck) (*CheckResult, error) {
	var piiResult *PiiDetectionResult = nil

	if check.ScanPii {
		piiRe, err := m.retrievePiiScore(check.Prompt)
		piiResult = piiRe

		if err != nil {
			return nil, err
		}
	}

	return &CheckResult{piiResult}, nil
}

func (m *Moat) retrievePiiScore(basePrompt string) (*PiiDetectionResult, error) {
	piiResult, err := m.piiScanner.Scan(basePrompt)

	if err != nil {
		fmt.Printf("Error scanning for PII %v\n", err)
		return nil, err
	}

	return &PiiDetectionResult{ContainsPii: piiResult.ContainingPii}, nil
}

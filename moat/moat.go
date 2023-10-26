// Package moat provides the first layer of defence for applications, by checking for attempted 'jailbreaks', and
// to scan a given text prompt for Personally Identifiable Information (PII) using an external PII scanner.
//
// Jailbreaking is the process of attempting to subvert the application's security defences, and is a commonly used
// to bypass ChatGPTs filters and generate offensive content. In this context, we are looking to see if a user is inputting
// user data to try and force the application or the LLM to generate a response that does something it shouldn't.
//
// For more information on PII scanning, see the pii and pii_aws packages documentation.
package moat

import (
	"fmt"
	"github.com/safetorun/PromptDefender/badwords"
	"github.com/safetorun/PromptDefender/pii"
)

type Moat struct {
	piiScanner    pii.Scanner
	badWordsCheck *badwords.BadWords
}

type PromptToCheck struct {
	Prompt  string
	ScanPii bool
}

type PiiDetectionResult struct {
	ContainsPii bool
}

type CheckResult struct {
	PiiResult        *PiiDetectionResult
	ContainsBadWords bool
}

func New(piiScanner pii.Scanner, badWordsCheck *badwords.BadWords) *Moat {
	return &Moat{piiScanner: piiScanner, badWordsCheck: badWordsCheck}
}

func (m *Moat) CheckMoat(check PromptToCheck) (*CheckResult, error) {
	var piiResult *PiiDetectionResult = nil

	containsBadWords, err := m.badWordsCheck.CheckPromptContainsBadWords(check.Prompt)

	if err != nil || containsBadWords == nil {
		return nil, fmt.Errorf("error checking bad words: %v", err)
	}

	if check.ScanPii {
		piiRe, err := m.retrievePiiScore(check.Prompt)
		piiResult = piiRe

		if err != nil {
			return nil, err
		}
	}

	return &CheckResult{PiiResult: piiResult, ContainsBadWords: *containsBadWords}, nil
}

func (m *Moat) retrievePiiScore(basePrompt string) (*PiiDetectionResult, error) {
	piiResult, err := m.piiScanner.Scan(basePrompt)

	if err != nil {
		fmt.Printf("Error scanning for PII %v\n", err)
		return nil, err
	}

	return &PiiDetectionResult{ContainsPii: piiResult.ContainingPii}, nil
}

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
	PiiScanner         pii.Scanner
	BadWordsCheck      *badwords.BadWords
	XmlEscapingScanner XmlEscapingScanner
}

type PromptToCheck struct {
	Prompt           string
	ScanPii          bool
	XmlTagToCheckFor *string
}

type PiiDetectionResult struct {
	ContainsPii bool
}

type CheckResult struct {
	PiiResult        *PiiDetectionResult
	ContainsBadWords bool
	XmlScannerResult *XmlEscapingDetectionResult
}

type MoatOpts func(*Moat) error

func New(opts ...MoatOpts) (*Moat, error) {
	m := &Moat{}

	for _, opt := range opts {
		err := opt(m)
		if err != nil {
			return nil, err
		}
	}

	return m, nil
}

func (m *Moat) CheckMoat(check PromptToCheck) (*CheckResult, error) {
	var piiResult *PiiDetectionResult = nil
	var xmlResult *XmlEscapingDetectionResult = nil

	containsBadWords, err := m.BadWordsCheck.CheckPromptContainsBadWords(check.Prompt)

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

	if check.XmlTagToCheckFor != nil {
		xmlResultInner, err := m.XmlEscapingScanner.Scan(check.Prompt, *check.XmlTagToCheckFor)
		if err != nil {
			return nil, err
		}

		xmlResult = xmlResultInner
	}

	return &CheckResult{PiiResult: piiResult, ContainsBadWords: *containsBadWords, XmlScannerResult: xmlResult}, nil
}

func (m *Moat) retrievePiiScore(basePrompt string) (*PiiDetectionResult, error) {
	piiResult, err := m.PiiScanner.Scan(basePrompt)

	if err != nil {
		fmt.Printf("Error scanning for PII %v\n", err)
		return nil, err
	}

	return &PiiDetectionResult{ContainsPii: piiResult.ContainingPii}, nil
}

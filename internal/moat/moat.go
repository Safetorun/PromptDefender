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
	"github.com/safetorun/PromptDefender/tracer"
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

func (m *Moat) CheckMoat(check PromptToCheck, t tracer.Tracer) (*CheckResult, error) {

	var piiResult *PiiDetectionResult = nil
	var xmlResult *XmlEscapingDetectionResult = nil

	containsBadWords, err := m.checkPromptContainsBadwords(check, t)

	if err != nil {
		return nil, err
	}

	if check.ScanPii {
		piiRe, err := m.checkPromptForPii(check, t)
		piiResult = piiRe

		if err != nil {
			return nil, err
		}
	}

	if check.XmlTagToCheckFor != nil {
		xmlResultInner, err := m.checkForXmlEscaping(check, t)
		if err != nil {
			return nil, err
		}

		xmlResult = xmlResultInner

	}

	return &CheckResult{PiiResult: piiResult, ContainsBadWords: *containsBadWords, XmlScannerResult: xmlResult}, nil
}

func (m *Moat) checkForXmlEscaping(check PromptToCheck, t tracer.Tracer) (*XmlEscapingDetectionResult, error) {
	wrappedMethod := tracer.TracerGenericsWrapper2[string, string, *XmlEscapingDetectionResult](m.XmlEscapingScanner.Scan)
	xmlResultInner, err := t.TraceDecorator(wrappedMethod, "xml_check")(check.Prompt)
	return xmlResultInner.(*XmlEscapingDetectionResult), err
}

func (m *Moat) checkPromptForPii(check PromptToCheck, t tracer.Tracer) (*PiiDetectionResult, error) {
	wrappedMethod := tracer.TracerGenericsWrapper[string, *PiiDetectionResult](m.retrievePiiScore)
	piiRe, err := t.TraceDecorator(wrappedMethod, "scanning_pii")(check.Prompt)
	return piiRe.(*PiiDetectionResult), err
}

func (m *Moat) checkPromptContainsBadwords(check PromptToCheck, t tracer.Tracer) (*bool, error) {
	// Wrap the method in a tracer
	wrappedMethod := tracer.TracerGenericsWrapper[string, *bool](m.BadWordsCheck.CheckPromptContainsBadWords)

	// Execute m.BadWordsCheck.CheckPromptContainsBadWords with the wrapped method
	containsBadWords, err := t.TraceDecorator(wrappedMethod, "scanning_bad_words")(check.Prompt)

	if err != nil || containsBadWords == nil {
		return nil, fmt.Errorf("error checking bad words: %v", err)
	}

	return containsBadWords.(*bool), nil
}

func (m *Moat) retrievePiiScore(basePrompt string) (*PiiDetectionResult, error) {
	piiResult, err := m.PiiScanner.Scan(basePrompt)

	if err != nil {
		fmt.Printf("Error scanning for PII %v\n", err)
		return nil, err
	}

	return &PiiDetectionResult{ContainsPii: piiResult.ContainingPii}, nil
}

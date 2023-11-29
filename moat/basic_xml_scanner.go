package moat

import "strings"

type BasicXmlScanner struct {
}

func NewBasicXmlEscapingScaner() XmlEscapingScanner {
	return &BasicXmlScanner{}
}

func (b BasicXmlScanner) Scan(textToScan string, tagToScanFor string) (*XmlEscapingDetectionResult, error) {

	if tagToScanFor == "" {
		return &XmlEscapingDetectionResult{
			ContainsXmlEscaping: false,
		}, nil
	}

	re := XmlEscapingDetectionResult{
		ContainsXmlEscaping: strings.Contains(textToScan, tagToScanFor),
	}

	return &re, nil
}

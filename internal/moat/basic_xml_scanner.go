package moat

import "strings"

type BasicXmlScanner struct {
}

func NewBasicXmlEscapingScaner() XmlEscapingScanner {
	return &BasicXmlScanner{}
}

func (b BasicXmlScanner) Scan(textToScan string, tagContentToScanFor string) (*XmlEscapingDetectionResult, error) {

	if tagContentToScanFor == "" {
		return &XmlEscapingDetectionResult{
			ContainsXmlEscaping: false,
		}, nil
	}

	tagToScanFor := "<" + tagContentToScanFor + ">"
	otherTagToScanFor := "</" + tagContentToScanFor + ">"

	re := XmlEscapingDetectionResult{
		ContainsXmlEscaping: strings.Contains(textToScan, tagToScanFor) ||
			strings.Contains(textToScan, otherTagToScanFor),
	}

	return &re, nil
}

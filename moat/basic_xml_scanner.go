package moat

type BasicXmlScanner struct {
}

func NewBasicXmlEscapingScaner() XmlEscapingScanner {
	return &BasicXmlScanner{}
}

func (b BasicXmlScanner) Scan(textToScan string, tagToScanFor string) (*XmlEscapingDetectionResult, error) {
	return nil, nil
}

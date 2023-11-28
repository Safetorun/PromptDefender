package moat

type XmlEscapingDetectionResult struct {
}

type XmlEscapingScanner interface {
	Scan(textToScan string, tagToScanFor string) (*XmlEscapingDetectionResult, error)
}

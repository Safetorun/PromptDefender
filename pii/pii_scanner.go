package pii

type PiiScanResult struct {
    ContainingPii bool
}

type PiiScanner interface {
    Scan(pii string) (PiiScanResult, error)
}

package pii_aws

import (
	"github.com/safetorun/PromptShield/pii"
)

type AwsPIIScanner struct {
}

func New() AwsPIIScanner {
	return AwsPIIScanner{}
}

func (s AwsPIIScanner) Scan(pii string) (pii.PiiScanResult, error) {
	return pii.PiiScanResult{ContainingPii: false}, nil
}

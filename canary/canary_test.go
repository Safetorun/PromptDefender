package canary

import (
	"testing"
)

func TestNew(t *testing.T) {
	checker := New(Basic, "test")
	if checker.Mode != Basic || checker.Canary != "test" {
		t.Errorf("New(Basic, \"test\") = %v; want Mode=Basic and Canary=test", checker)
	}
}

func TestCheckBasic(t *testing.T) {
	checker := New(Basic, "testCanary")

	result, err := checker.Check("testCanary")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !result.Detected {
		t.Errorf("Canary not detected, got: %v, want: true", result.Detected)
	}

	result, err = checker.Check("noMatch")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result.Detected {
		t.Errorf("Canary detected, got: %v, want: false", result.Detected)
	}
}

func TestCheckNoneMode(t *testing.T) {
	checker := New(None, "testCanary")

	result, err := checker.Check("testCanary")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result.Detected {
		t.Errorf("Canary detected, got: %v, want: false", result.Detected)
	}
}

func TestCheckUnsupportedMode(t *testing.T) {
	checker := New(Intermediate, "testCanary")

	_, err := checker.Check("testCanary")
	if err == nil || err.Error() != "mode not implemented" {
		t.Errorf("Unexpected error, got: %v, want: Mode not implemented", err)
	}
}

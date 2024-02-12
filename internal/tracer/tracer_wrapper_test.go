package tracer

import (
	"testing"
)

func TestTracerWrapperGenerics(t *testing.T) {
	expectedOutput := "Hello world"
	input := "Hello"
	testFunc := func(input string) (string, error) {
		return input + " world", nil
	}

	response, err := TracerGenericsWrapper(testFunc)(input)

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if response != expectedOutput {
		t.Errorf("Expected %v but got %v", expectedOutput, response)
	}
}

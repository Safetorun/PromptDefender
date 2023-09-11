package game

import (
	"strings"
	"testing"
)

func TestNewPromptBuilder(t *testing.T) {
	pb := NewPromptBuilder()

	if pb == nil || !strings.Contains(pb.Build(), BasePrompt) {
		t.Errorf("NewPromptBuilder() failed, expected base prompt but got %v", pb)
	}
}

func TestThatPrefixAddsAPrefix(t *testing.T) {
	prefix := "some prefix"
	pb := NewPromptBuilder()
	pb.AddPrefixPromptDefense(prefix)

	if !strings.Contains(pb.Build(), prefix) {
		t.Errorf("AddPrefixPromptDefense() failed, prefix not found in prompt")
	}
}

func TestThatSuffixAddsASuffix(t *testing.T) {
	suffix := "some suffix"
	pb := NewPromptBuilder()
	pb.AddSuffixPromptDefense(suffix)

	if !strings.Contains(pb.Build(), suffix) {
		t.Errorf("AddSuffixPromptDefense() failed, suffix not found in prompt")
	}
}

func TestThatPrefixAndSuffixCanBeAdded(t *testing.T) {
	prefix := "some prefix"
	suffix := "some suffix"
	pb := NewPromptBuilder()
	pb.AddPrefixPromptDefense(prefix).AddSuffixPromptDefense(suffix)

	if !strings.Contains(pb.Build(), prefix) || !strings.Contains(pb.Build(), suffix) {
		t.Errorf("AddPrefixPromptDefense() and AddSuffixPromptDefense() failed, prefix or suffix not found in prompt")
	}
}

func TestAddLoggedInUser(t *testing.T) {
	userId := "user123"
	pb := NewPromptBuilder()
	pb.AddLoggedInUser(userId)

	if !strings.Contains(pb.Build(), userId) {
		t.Errorf("AddLoggedInUser() failed, userId not found in prompt")
	}
}

func TestAddContext(t *testing.T) {
	context := "some context"
	pb := NewPromptBuilder()
	pb.AddContext(context)

	if !strings.Contains(pb.Build(), context) {
		t.Errorf("AddContext() failed, context not found in prompt")
	}
}

func TestBuild(t *testing.T) {
	userId := "user123"
	context := "some context"

	pb := NewPromptBuilder()
	pb.AddLoggedInUser(userId).AddContext(context)

	finalPrompt := pb.Build()

	if !strings.Contains(finalPrompt, userId) || !strings.Contains(finalPrompt, context) || !strings.Contains(finalPrompt, BasePrompt) {
		t.Errorf("Build() failed, prompt built incorrectly")
	}
}

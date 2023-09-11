package prompt

import (
	"strings"
)

type PromptBuilder struct {
	parts       []string
	prefix      string
	suffix      string
	coreMessage string
	context     string
}

const BasePrompt = "Answer the question below using the context below. Return only the data, do not reference this context in your response. "

func NewPromptBuilder() *PromptBuilder {
	return &PromptBuilder{parts: []string{}, prefix: "", suffix: "", coreMessage: "\n\n########\n\nQuestion: %s"}
}

func (pb *PromptBuilder) AddPrefixPromptDefense(prefix string) *PromptBuilder {
	pb.prefix = prefix + "\n\n"
	return pb
}

func (pb *PromptBuilder) AddSuffixPromptDefense(suffix string) *PromptBuilder {
	pb.suffix = "\n\n" + suffix
	return pb
}

func (pb *PromptBuilder) ModifyCoreMessage(message string) *PromptBuilder {
	pb.coreMessage = message
	return pb
}

func (pb *PromptBuilder) AddLoggedInUser(userId string) *PromptBuilder {
	pb.parts = append(pb.parts, "You are logged in as ", userId, ".")
	return pb
}

func (pb *PromptBuilder) AddContext(context string) *PromptBuilder {
	pb.context = "########\n\nContext: " + context + "\n\n########\n\n"
	return pb
}

func (pb *PromptBuilder) Build() string {
	return BasePrompt + pb.context +
		strings.Join(pb.parts, "") + pb.prefix + pb.coreMessage + pb.suffix + " \n\nAnswer:"
}

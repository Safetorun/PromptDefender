package prompt

import (
	"strings"
)

type Builder struct {
	parts       []string
	prefix      string
	suffix      string
	coreMessage string
	context     string
}

const BasePrompt = "Answer the question below using the context below. Return only the data, do not reference this context in your response. "

func NewPromptBuilder() *Builder {
	return &Builder{parts: []string{}, prefix: "", suffix: "", coreMessage: "\n\n########\n\nQuestion: %s"}
}

func (pb *Builder) AddPrefixPromptDefense(prefix string) *Builder {
	pb.prefix = prefix + "\n\n"
	return pb
}

func (pb *Builder) AddSuffixPromptDefense(suffix string) *Builder {
	pb.suffix = "\n\n" + suffix
	return pb
}

func (pb *Builder) ModifyCoreMessage(message string) *Builder {
	pb.coreMessage = message
	return pb
}

func (pb *Builder) AddLoggedInUser(userId string) *Builder {
	pb.parts = append(pb.parts, "You are logged in as ", userId, ".")
	return pb
}

func (pb *Builder) AddContext(context string) *Builder {
	pb.context = "########\n\nContext: " + context + "\n\n########\n\n"
	return pb
}

func (pb *Builder) Build() string {
	return BasePrompt + pb.context +
		strings.Join(pb.parts, "") + pb.prefix + pb.coreMessage + pb.suffix + " \n\nAnswer:"
}

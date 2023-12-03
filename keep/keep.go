package keep

import (
	"github.com/safetorun/PromptDefender/aiprompt"
	"github.com/safetorun/PromptDefender/prompt"
	"math/rand"
	"time"
)

type Callback func(prompt string, newPrompt string) error

type KeepOption func(*Keep)

type Keep struct {
	openAi   aiprompt.RemoteAIChecker
	Callback *Callback
}

type StartingPrompt struct {
	Prompt       string
	RandomiseTag bool
}

type NewPrompt struct {
	NewPrompt string
	Tag       string
}

func New(aiPrompt aiprompt.RemoteAIChecker, options ...KeepOption) *Keep {
	k := &Keep{
		openAi: aiPrompt,
	}

	for _, opt := range options {
		opt(k)
	}

	return k
}

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func (k *Keep) BuildKeep(startingPrompt StartingPrompt) (*NewPrompt, error) {

	tag := "user_input"

	if startingPrompt.RandomiseTag {
		tag = generateRandomString(10)
	}

	builtPrompt := prompt.SmartPrompt(prompt.SmartPromptRequest{BasePrompt: startingPrompt.Prompt, XmlTagName: tag})

	response, err := k.openAi.CheckAI(builtPrompt)

	if err != nil {
		return nil, err
	}

	if k.Callback != nil {
		err := (*k.Callback)(startingPrompt.Prompt, *response)
		if err != nil {
			return nil, err
		}
	}

	return &NewPrompt{NewPrompt: *response, Tag: tag}, nil
}

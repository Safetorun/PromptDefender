package embeddings

import (
	"context"
	"github.com/sashabaranov/go-openai"
)

type OpenAI struct {
	ApiKey string
}

func (o OpenAI) CreateEmbeddings(inputText string) (*EmbeddingValue, error) {
	client := openai.NewClient(o.ApiKey)
	response, err := client.CreateEmbeddings(context.Background(), openai.EmbeddingRequest{
		Input: []string{inputText},
		Model: openai.AdaEmbeddingV2,
	})

	if err != nil {
		return nil, err
	}

	return &EmbeddingValue{response.Data[0].Embedding}, nil
}

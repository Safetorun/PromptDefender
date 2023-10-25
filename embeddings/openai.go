package embeddings

import (
	"context"
	"strconv"
	"strings"

	openai "github.com/sashabaranov/go-openai"
)

type OpenAI struct {
	ApiKey string
}

func New(apiKey string) *OpenAI {
	return &OpenAI{ApiKey: apiKey}
}

func convertFloat32ToFloat64(float32Array []float32) []float64 {
	float64Array := make([]float64, len(float32Array))
	for i, v := range float32Array {
		float64Array[i] = float64(v)
	}
	return float64Array
}

func stringToFloatArray(str string) ([]float64, error) {
	strs := strings.Split(str, ",")
	floats := make([]float64, len(strs))
	for i, s := range strs {
		f, err := strconv.ParseFloat(strings.TrimSpace(s), 64)
		if err != nil {
			return nil, err
		}
		floats[i] = f
	}
	return floats, nil
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

	return &EmbeddingValue{convertFloat32ToFloat64(response.Data[0].Embedding)}, nil
}

func (o OpenAI) RetrieveBadwordsEmbeddings() (*[]EmbeddingValue, error) {
	var embeddingValues []EmbeddingValue

	for _, badword := range GetBadwords() {
		embeddingValues = append(embeddingValues, EmbeddingValue{EmbeddingValue: badword})
	}

	return &embeddingValues, nil
}

package badwords_embeddings

import (
	"github.com/drewlanenga/govector"
	"github.com/safetorun/PromptDefender/badwords"
	"github.com/safetorun/PromptDefender/embeddings"
)

type BadwordsEmbeddings struct {
	embeddings embeddings.Embeddings
}

func New(embeddings embeddings.Embeddings) BadwordsEmbeddings {
	return BadwordsEmbeddings{embeddings: embeddings}
}

func (bw BadwordsEmbeddings) GetClosestMatch(prompt string) (*badwords.ClosestMatchScore, error) {
	promptEmbeddingsValue, err := bw.embeddings.CreateEmbeddings(prompt)

	if err != nil {
		return nil, err
	}

	embeddedBadwords, err := bw.embeddings.RetrieveBadwordsEmbeddings()

	if err != nil {
		return nil, err
	}

	var lowestScore float64 = 1.0

	for _, embeddings := range *embeddedBadwords {
		newScore := cosineSimilarity(promptEmbeddingsValue.EmbeddingValue, embeddings.EmbeddingValue)

		if newScore < lowestScore {
			lowestScore = newScore
		}
	}

	return &badwords.ClosestMatchScore{Score: lowestScore}, nil
}

func cosineSimilarity(a, b []float64) float64 {
	vector1, _ := govector.AsVector(a)
	vector2, _ := govector.AsVector(b)
	cosineSimilarity, _ := govector.Cosine(vector1, vector2)
	cosineDistance := 1 - cosineSimilarity
	return cosineDistance
}

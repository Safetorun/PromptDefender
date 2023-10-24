package badwords_embeddings

import (
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
	err, promptEmbeddingsValue := bw.embeddings.CreateEmbeddings(prompt)

	if err != nil {
		return nil, err
	}

	// TODO: implement with embeddings probably loaded from file.
}

package embeddings

type EmbeddingValue struct {
	EmbeddingValue []float32
}

type Embeddings interface {
	CreateEmbeddings(string) (*EmbeddingValue, error)
}

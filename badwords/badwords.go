package badwords

type ClosestMatchScore struct {
	MatchDescription string
	Score            float64
}

type ClosestMatcher interface {
	GetClosestMatch(string) (*ClosestMatchScore, error)
}

type BadWords struct {
	matcher   ClosestMatcher
	threshold float64
}

type BadWordsOption func(*BadWords)

func New(matcher ClosestMatcher, opts ...BadWordsOption) *BadWords {
	badWords := &BadWords{matcher: matcher, threshold: 0.1}
	for _, opt := range opts {
		opt(badWords)
	}

	return badWords
}

func (badWords BadWords) CheckPromptContainsBadWords(prompt string) (*bool, error) {
	score, err := badWords.matcher.GetClosestMatch(prompt)

	if err != nil {
		return nil, err
	}

	containsBadWord := score.Score < badWords.threshold
	return &containsBadWord, nil

}

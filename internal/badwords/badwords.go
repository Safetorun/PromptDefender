package badwords

type MatchLevel int

const (
	ExactMatch MatchLevel = iota
	VeryClose
	Medium
	NoMatch
	TotallyDifferent
)

type ClosestMatchScore struct {
	MatchDescription string
	Score            MatchLevel
}

type ClosestMatcher interface {
	GetClosestMatch(string) (*ClosestMatchScore, error)
}

type BadWords struct {
	matcher   ClosestMatcher
	threshold MatchLevel
}

type BadWordsOption func(*BadWords)

func New(matcher ClosestMatcher, opts ...BadWordsOption) *BadWords {
	badWords := &BadWords{matcher: matcher, threshold: VeryClose}
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

package wall

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

type MatchLevel int

const (
	ExactMatch MatchLevel = iota
	VeryClose
	Medium
	NoMatch
	TotallyDifferent
)

const (
	ApiUrl = "https://api-inference.huggingface.co/models/deepset/deberta-v3-base-injection"
)

type InjectionResponse []struct {
	Label string  `json:"label"`
	Score float64 `json:"score"`
}

type RemoteApiCallerImpl struct {
	huggingfaceToken string // token to call the remote API
	huggingfaceUrl   string // URL of the remote API
	logger           *log.Logger
}

type Payload struct {
	Inputs string `json:"inputs"`
}

func NewRemoteApiCaller(huggingfaceToken string) RemoteApiCallerImpl {
	return RemoteApiCallerImpl{
		huggingfaceToken: huggingfaceToken,
		huggingfaceUrl:   ApiUrl,
		logger:           log.Default(),
	}
}

func (c *RemoteApiCallerImpl) Query(payload Payload) (*float64, error) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.huggingfaceUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.huggingfaceToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	println(string(body))
	var result []InjectionResponse

	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	for _, re := range result[0] {
		if re.Label == "INJECTION" {
			return &re.Score, nil
		}
	}

	return nil, errors.New(
		fmt.Sprintf("Could not find the label in the response. Response %+v", result),
	)
}

// CallRemoteApi calls the remote API and returns the injection score
func (r *RemoteApiCallerImpl) CallRemoteApi(prompt string) (MatchLevel, error) {
	response, err := r.Query(Payload{Inputs: prompt})

	r.logger.Println("Prompt is ", prompt, " Injection score is ", *response)

	if err != nil {
		return -1, err
	}

	return matchLevelForScore(*response), nil
}

func matchLevelForScore(score float64) MatchLevel {
	if score == 1.0 {
		return ExactMatch
	}
	if score > 0.9 {
		return VeryClose
	}
	if score > 0.5 {
		return Medium
	}

	return NoMatch
}

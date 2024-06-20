package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	BaseUselessFactsEndPoint = "https://uselessfacts.jsph.pl/api/v2/facts/random"
)

type UselessFacts struct {
	ID        string `json:"id"`
	Text      string `json:"text"`
	Source    string `json:"source"`
	SourceURL string `json:"source_url"`
	Language  string `json:"language"`
	Permalink string `json:"permalink"`
}

func (c *Client) GetUselessFact() (UselessFacts, error) {
	const errorString = "could not get uselessfact: %v"

	req, err := http.NewRequest("GET", BaseUselessFactsEndPoint, nil)
	if err != nil {
		return UselessFacts{}, fmt.Errorf(errorString, err)
	}

	res, err := c.client.Do(req)
	if err != nil {
		return UselessFacts{}, fmt.Errorf(errorString, err)
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		return UselessFacts{}, fmt.Errorf(errorString, fmt.Sprintf("unexpected status code: %d", res.StatusCode))
	}

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return UselessFacts{}, fmt.Errorf(errorString, err)
	}

	uselessFacts := UselessFacts{}
	if err := json.Unmarshal(dat, &uselessFacts); err != nil {
		return UselessFacts{}, fmt.Errorf(errorString, err)
	}
	return uselessFacts, nil
}

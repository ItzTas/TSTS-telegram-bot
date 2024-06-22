package client

import (
	"encoding/json"
	"fmt"
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
	dat, err := c.getURLData(BaseUselessFactsEndPoint, nil)
	if err != nil {
		return UselessFacts{}, fmt.Errorf("could not get uselessfact: %v", err)
	}

	uselessFacts := UselessFacts{}
	if err := json.Unmarshal(dat, &uselessFacts); err != nil {
		return UselessFacts{}, fmt.Errorf("could not get uselessfact: %v", err)
	}
	return uselessFacts, nil
}

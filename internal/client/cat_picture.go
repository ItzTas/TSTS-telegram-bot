package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

const (
	BaseCatImageURL = "https://api.thecatapi.com/v1/images/search"
)

type CatImage struct {
	ID     string `json:"id"`
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

func (c *Client) GetCatPictures(limit string) ([]CatImage, error) {
	h := http.Header{
		"x-api-key": []string{c.APIKeys["catAPIKey"]},
	}

	l, err := strconv.Atoi(limit)
	if err != nil || limit == "" {
		limit = "1"
	}

	if l > 12 {
		limit = "12"
	}

	url := BaseCatImageURL + "?limit=" + limit

	dat, err := c.getURLData(url, nil, h)
	if err != nil {
		return []CatImage{}, fmt.Errorf("could not get picture: %v", err)
	}

	catImages := make([]CatImage, 0)
	if err := json.Unmarshal(dat, &catImages); err != nil {
		return []CatImage{}, fmt.Errorf("could not get picture: %v", err)
	}

	return catImages, nil
}

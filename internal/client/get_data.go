package client

import (
	"io"
	"net/http"
)

func (c *Client) getURLData(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []byte{}, err
	}

	res, err := c.client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		return []byte{}, err
	}

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}

	return dat, err
}
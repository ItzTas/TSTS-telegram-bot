package client

import (
	"io"
	"net/http"
)

func (c *Client) getURLData(url string, body io.Reader, header http.Header) ([]byte, error) {
	req, err := http.NewRequest("GET", url, body)
	if err != nil {
		return []byte{}, err
	}

	req.Header = header

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
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

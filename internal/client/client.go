package client

import (
	"net/http"
	"time"
)

type Client struct {
	client  http.Client
	APIKeys map[string]string
}

func NewClient(timeout time.Duration, APIKeys map[string]string) Client {
	c := Client{
		client: http.Client{
			Timeout: timeout,
		},
		APIKeys: APIKeys,
	}

	return c
}

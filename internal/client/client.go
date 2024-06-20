package client

import (
	"net/http"
	"time"
)

type Client struct {
	client http.Client
}

func NewClient(timeout time.Duration) Client {
	c := Client{
		client: http.Client{
			Timeout: timeout,
		},
	}

	return c
}

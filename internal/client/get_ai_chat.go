package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	BaseAIEndpoint = "http://127.0.0.1:5000/chat"
)

func (c *Client) GetaiChat(content string) (string, error) {
	const errorString = "could not get chat: %v"
	type Payload struct {
		Content string `json:"content"`
	}

	payload := Payload{
		Content: content,
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf(errorString, err)
	}

	dat, err := c.getURLData(BaseAIEndpoint, bytes.NewBuffer(body), http.Header{})
	if err != nil {
		return "", fmt.Errorf(errorString, err)
	}
	resPayload := Payload{}
	if err := json.Unmarshal(dat, &resPayload); err != nil {
		return "", fmt.Errorf(errorString, err)
	}

	return resPayload.Content, nil
}

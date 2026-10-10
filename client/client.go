package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var (
	ErrInvalidJSON = errors.New("unable marshal to JSON")
)

type Request struct {
	Prompt string `json:"prompt"`
}

type Client struct {
	Endpoint string
}

func (c Client) Send(text string) error {
	req := Request{Prompt: text}
	j, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("encode request: %w", err)
	}
	request, err := http.NewRequest("POST", c.Endpoint, bytes.NewReader(j))
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}
	request.Header.Set("Content-Type", "application/json")
	httpClient := &http.Client{}
	response, err := httpClient.Do(request)
	if err != nil {
		return fmt.Errorf("send request: %w", err)
	}

	defer response.Body.Close()

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return fmt.Errorf("send request unexpected HTTP status %s", response.Status)
	}
	return nil
}

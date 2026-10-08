package client

import (
	"encoding/json"
	"errors"
	"net/http"
)

var (
	ErrInvalidJSON = errors.New("unable marshal to JSON")
)

type ServerConfig struct {
	endpoint string
}

type Message struct {
	text string
}

func (m Message) toJSON() ([]byte, error) {
	j, err := json.Marshal(m.text)
	if err != nil {
		return nil, ErrInvalidJSON
	}
	return j, nil
}

func Send(text string) {
	http.Get(ServerConfig)
}

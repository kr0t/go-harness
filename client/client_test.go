package client

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClientSendSuccess(t *testing.T) {
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodPost {
				t.Errorf("request method = %q, want POST", r.Method)
			}
			contentType := r.Header.Get("Content-Type")
			if contentType != "application/json" {
				t.Errorf(
					"Content-Type = %q, want application/json",
					contentType,
				)
			}
			var received Request

			if err := json.NewDecoder(r.Body).Decode(&received); err != nil {
				t.Errorf("cannot decode request JSON: %v", err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			if received.Prompt != "Explain RAG architecture" {
				t.Errorf(
					"request prompt = %q, want %q",
					received.Prompt,
					"Explain RAG architecture",
				)
			}
			w.WriteHeader(http.StatusOK)
		}),
	)
	t.Cleanup(server.Close)

	c := Client{Endpoint: server.URL}

	err := c.Send("Explain RAG architecture")

	if err != nil {
		t.Fatalf("Send() returned unexpected error: %v", err)
	}

}

func TestClientSendHTTPError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		},
	))
	t.Cleanup(server.Close)

	c := Client{Endpoint: server.URL}

	err := c.Send("Explain RAG architecture")

	if err == nil {
		t.Fatal("Send() returned nil, want error for HTTP 500")
	}
}

func TestClientSendTransportError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		},
	))
	t.Cleanup(server.Close)

	url := server.URL

	server.Close()
	c := Client{Endpoint: url}

	err := c.Send("Explain RAG architecture")

	if err == nil {
		t.Fatal("Send() returned nil, want transport error")
	}
}

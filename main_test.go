package main

import (
	"errors"
	"promtctl/prompt"
	"testing"
)

func TestProcessPrompt(t *testing.T) {
	tests := []struct {
		name    string
		text    string
		wantErr error
	}{
		{
			name:    "empty prompt",
			text:    "",
			wantErr: prompt.ErrEmptyPrompt,
		},
		{
			name:    "too short prompt",
			text:    "Explain RAG",
			wantErr: prompt.ErrPromptTooShort,
		},
		{
			name:    "valid prompt",
			text:    "Explain RAG architecture",
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := prompt.Prompt{Text: tt.text}
			err := processPrompt(p)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("processPrompt() error = %v, want %v", err, tt.wantErr)
			}
		})
	}
}

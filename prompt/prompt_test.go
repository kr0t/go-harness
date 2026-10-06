package prompt_test

import (
	"errors"
	"promtctl/prompt"
	"testing"
)

func TestPromptValidate(t *testing.T) {
	tests := []struct {
		name    string
		text    string
		wantErr error
	}{
		{
			name:    "valid ASCII prompt",
			text:    "Explain RAG architecture",
			wantErr: nil,
		},
		{
			name:    "two words",
			text:    "Explain RAG",
			wantErr: prompt.ErrPromptTooShort,
		},
		{
			name:    "empty prompt",
			text:    "",
			wantErr: prompt.ErrEmptyPrompt,
		},
		{
			name:    "only spaces",
			text:    "      ",
			wantErr: prompt.ErrEmptyPrompt,
		},
		{
			name:    "valid Unicode prompt",
			text:    "Объясни архитектуру RAG",
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := prompt.Prompt{Text: tt.text}

			err := p.ValidatePrompt()

			if !errors.Is(err, tt.wantErr) {
				t.Errorf("ValidatePrompt() error = %v, want %v", err, tt.wantErr)
			}
		})
	}
}

func TestPromptStats(t *testing.T) {
	tests := []struct {
		name           string
		text           string
		wantWords      int
		wantCharacters int
		wantValid      bool
	}{
		{
			name:           "ASCII prompt",
			text:           "Explain RAG architecture",
			wantWords:      3,
			wantCharacters: 24,
			wantValid:      true,
		},
		{
			name:           "Unicode prompt",
			text:           "Объясни архитектуру RAG",
			wantWords:      3,
			wantCharacters: 23,
			wantValid:      true,
		},
		{
			name:           "empty prompt",
			text:           "",
			wantWords:      0,
			wantCharacters: 0,
			wantValid:      false,
		},
		{
			name:           "only spaces",
			text:           "   ",
			wantWords:      0,
			wantCharacters: 3,
			wantValid:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := prompt.Prompt{Text: tt.text}

			got := p.Stats()

			if got.Prompt != tt.text {
				t.Errorf("Stats().Prompt = %q, want %q", got.Prompt, tt.text)
			}

			if got.Words != tt.wantWords {
				t.Errorf("Stats().Words = %d, want %d", got.Words, tt.wantWords)
			}

			if got.Characters != tt.wantCharacters {
				t.Errorf("Stats().Characters = %d, want %d", got.Characters, tt.wantCharacters)
			}

			if got.Valid != tt.wantValid {
				t.Errorf("Stats().Valid = %v, want %v", got.Valid, tt.wantValid)
			}

		})
	}
}

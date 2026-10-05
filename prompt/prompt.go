package prompt

import (
	"errors"
	"strings"
	"unicode/utf8"
)

var (
	ErrEmptyPrompt    = errors.New("prompt empty")
	ErrPromptTooShort = errors.New("prompt too short")
)

type Prompt struct {
	Text string
}

type PromptStats struct {
	Prompt     string
	Words      int
	Characters int
	Valid      bool
}

func (p Prompt) Stats() PromptStats {
	valid := p.ValidatePrompt()
	return PromptStats{
		Prompt:     p.Text,
		Words:      len(strings.Fields(p.Text)),
		Characters: utf8.RuneCountInString(p.Text),
		Valid:      valid == nil,
	}

}

func (p Prompt) ValidatePrompt() error {
	words := strings.Fields(p.Text)
	if len(words) == 0 {
		return ErrEmptyPrompt
	}
	if len(words) < 3 {
		return ErrPromptTooShort
	}
	return nil

}

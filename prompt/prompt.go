package prompt

import (
	"errors"
	"strings"
	"unicode/utf8"
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
	if len(strings.Fields(p.Text)) <= 2 {
		return errors.New("prompt must be at least 3 words")
	}
	return nil
}

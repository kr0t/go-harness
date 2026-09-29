package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

type Prompt struct {
	text string
}

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("error: command required")
		return
	} else if len(args) < 3 && !checkCommand(args[1]) {
		fmt.Println("Unknown command")
		return
	} else if len(args) < 3 {
		fmt.Println("error: prompt required")
		return
	}

	command := args[1]
	prompt := Prompt{args[2]}

	switch command {
	case "validate":
		prompt.validatePrompt()
	case "stats":
		prompt.promptStats()
	default:
		fmt.Println("Unknown command")
	}
}

func (p Prompt) promptStats() {
	fmt.Println("Prompt: ", p.text)
	fmt.Println("Words: ", len(strings.Fields(p.text)))
	fmt.Println("Characters: ", utf8.RuneCountInString(p.text))
}

func checkCommand(command string) bool {
	return command == "validate" || command == "stats"
}

func (p Prompt) validatePrompt() bool {
	if len(strings.Fields(p.text)) <= 2 {
		fmt.Println("error: prompt must be at least 3 words.")
		return false
	}
	fmt.Println("prompt is valid")
	return true
}
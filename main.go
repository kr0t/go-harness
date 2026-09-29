package main

import (
	"errors"
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
		//prompt.validatePrompt()
		err := prompt.validatePrompt()
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("prompt is valid")
		}
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

func (p Prompt) validatePrompt() error {
	if len(strings.Fields(p.text)) <= 2 {
		//fmt.Println("error: prompt must be at least 3 words.")
		return errors.New("prompt must be at least 3 words")
	}
	//fmt.Println("prompt is valid")
	return nil
}

package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

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
	switch command {
	case "validate":
		validatePrompt(args[2])
	case "stats":
		promptStats(args[2])
	default:
		fmt.Println("Unknown command")
	}
}

func validatePrompt(prompt string) {

	if len(strings.Fields(prompt)) <= 2 {
		fmt.Println("error: prompt must be at least 3 words.")
	} else {
		fmt.Println("prompt is valid")
	}

}

func promptStats(prompt string) {
	fmt.Println("Prompt: ", prompt)
	fmt.Println("Words: ", len(strings.Fields(prompt)))
	fmt.Println("Characters: ", utf8.RuneCountInString(prompt))
}

func checkCommand(command string) bool {
	return command == "validate" || command == "stats"
}

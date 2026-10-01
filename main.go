package main

import (
	"fmt"
	"os"
	"promtctl/prompt"
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
	p := prompt.Prompt{Text: args[2]}

	switch command {
	case "validate":
		err := p.ValidatePrompt()
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("prompt is valid")
		}
	case "stats":
		s := p.Stats()
		fmt.Println("Prompt: ", s.Prompt)
		fmt.Println("Characters: ", s.Characters)
		fmt.Println("Words: ", s.Words)
		fmt.Println("Valid: ", s.Valid)
	default:
		fmt.Println("Unknown command!")
	}
}

func checkCommand(command string) bool {
	return command == "validate" || command == "stats"
}

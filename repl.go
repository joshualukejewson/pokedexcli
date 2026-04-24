package main

import (
	"strings"
)

type clientCommand struct {
	name        string
	description string
	callback    func() error
}

func CleanInput(text string) []string {
	lowerStripped := strings.ToLower(strings.TrimSpace(text))
	cleanedInput := strings.Split(lowerStripped, " ")
	return cleanedInput
}

func getCommands() map[string]clientCommand {
	return map[string]clientCommand{
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
}

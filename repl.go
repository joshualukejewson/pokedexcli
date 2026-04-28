package main

import (
	"pokedexcli/api"
	"strings"
)

type clientCommand struct {
	name        string
	description string
	callback    func(*api.Config) error
}

func CleanInput(text string) []string {
	lowerStripped := strings.ToLower(strings.TrimSpace(text))
	cleanedInput := strings.Split(lowerStripped, " ")
	return cleanedInput
}

func GetCommands() map[string]clientCommand {
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
		"map": {
			name:        "map",
			description: "Print a list of the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Print a list of the previous 20 locations",
			callback:    commandMapBack,
		},
	}
}

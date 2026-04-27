package main

import (
	"fmt"
	"os"
	"pokedexcli/api"
)

func commandExit(cfg *api.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *api.Config) error {
	fmt.Print("Welcome to the Pokedex!\nUsage:\n\n")
	for _, cliCommand := range GetCommands() {
		fmt.Printf("%s: %s\n", cliCommand.name, cliCommand.description)
	}
	return nil
}

func commandMap(cfg *api.Config) error {
	for _, location := range cfg.Results {
		fmt.Printf("%s\n", location.Name)
	}
	return nil
}

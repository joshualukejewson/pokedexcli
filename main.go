package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedexcli/api"
)

const LOCATION_API_URL string = "https://pokeapi.co/api/v2/location-area"

func main() {

	cfg := api.Config{
		Next: LOCATION_API_URL,
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Pokedex > ")
	for scanner.Scan() {
		line := scanner.Text()
		args := CleanInput(line)
		cmd := args[0]
		cliCmd, ok := GetCommands()[cmd]
		if ok {
			if err := cliCmd.callback(&cfg); err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		} else {
			fmt.Println("Unknown command")
		}
		fmt.Print("Pokedex > ")
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading stdio:", err)
	}
}

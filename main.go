package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedexcli/api"
)

func main() {

	cfg, err := api.FillConfig()
	if err != nil {
		fmt.Printf("Error parsing config: %v", err)
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

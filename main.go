package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Pokedex > ")
	for scanner.Scan() {
		line := scanner.Text()
		args := CleanInput(line)
		cmd := args[0]
		cliCmd, ok := getCommands()[cmd]
		if ok {
			if err := cliCmd.callback(); err != nil {
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

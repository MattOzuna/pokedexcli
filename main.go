package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Pokedex > ")

	commands := map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}

	for scanner.Scan() {
		input := scanner.Text()

		command, ok := commands[input]
		if !ok {
			fmt.Println("Invalid command")
		} else {
			command.callback()
		}
		fmt.Print("Pokedex > ")
	}
}

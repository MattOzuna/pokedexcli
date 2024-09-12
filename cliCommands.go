package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func commandHelp(c *config) error {
	fmt.Print(`
Welcome to the Pokedex!
Usage:

`)
	for k, v := range getCommands() {
		fmt.Printf("%s: %s\n", k, v.description)
	}
	fmt.Println("")
	return nil
}

func commandExit(c *config) error {
	fmt.Println("goodbye!")
	os.Exit(0)
	return nil
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
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
		"map": {
			name:        "map",
			description: "Get next 20 locations",
			callback:    mapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get last 20 locations",
			callback:    mapb,
		},
	}
}

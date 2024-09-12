package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Pokedex > ")

	commands := getCommands()

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

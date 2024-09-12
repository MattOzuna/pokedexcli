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
	config := intialize()

	for scanner.Scan() {
		input := scanner.Text()

		command, ok := commands[input]
		if !ok {
			fmt.Println("Invalid command")
		} else {
			err := command.callback(&config)
			if err != nil {
				fmt.Println(err)
			}
		}
		fmt.Print("Pokedex > ")
	}
}

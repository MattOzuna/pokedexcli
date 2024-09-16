package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Pokedex > ")

	commands := getCommands()
	config := intialize()

	for scanner.Scan() {
		input := scanner.Text()
		input = strings.ToLower(input)
		inputs := strings.Fields(input)
		if len(inputs) < 2 {
			inputs = append(inputs, "")
		}

		command, ok := commands[inputs[0]]
		if !ok {
			fmt.Println("Invalid command")
		} else {
			err := command.callback(&config, inputs[1])
			if err != nil {
				fmt.Println(err)
			}
		}
		fmt.Print("Pokedex > ")
	}
}

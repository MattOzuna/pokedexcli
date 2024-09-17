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

	config := intialize()

	for scanner.Scan() {
		input := scanner.Text()
		input = strings.ToLower(input)
		words := strings.Fields(input)

		commandName := words[0]
		args := make([]string, len(words)-1)
		if len(words) > 1 {
			args = words[1:]
		}

		command, ok := getCommands()[commandName]
		if !ok {
			fmt.Println("Invalid command")
		} else {
			err := command.callback(&config, args...)
			if err != nil {
				fmt.Println(err)
			}
		}
		fmt.Print("Pokedex > ")
	}
}

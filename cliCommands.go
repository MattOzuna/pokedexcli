package main

import (
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandHelp() error {
	fmt.Println("help")
	return nil
}

func commandExit() error {
	fmt.Println("goodbye!")
	os.Exit(0)
	return nil
}

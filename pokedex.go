package main

import "fmt"

func pokedex(c *config, input ...string) error {
	pokedex := c.Pokedex

	if len(pokedex) == 0 {
		fmt.Println("You haven't caught any pokemon yet.")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for name := range pokedex {
		fmt.Printf(" -%v\n", name)
	}
	return nil
}

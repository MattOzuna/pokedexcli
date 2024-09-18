package main

import "fmt"

func inspect(c *config, input ...string) error {
	pokemon, isCaught := c.Pokedex[input[0]]
	if !isCaught {
		fmt.Println("you have not caught that pokemon")
		return nil
	}
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf(" -%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, Type := range pokemon.Types {
		fmt.Printf(" -%v\n", Type.Type.Name)
	}
	return nil
}

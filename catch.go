package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

func catch(c *config, input ...string) error {
	reqURL := "https://pokeapi.co/api/v2/pokemon/" + input[0]

	var data Pokemon
	res, err := http.Get(reqURL)
	if err != nil {
		return fmt.Errorf("issue reaching API: %v", err)
	}
	defer res.Body.Close()

	resData, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("issue Reading from response body: %v", err)
	}

	if err := json.Unmarshal(resData, &data); err != nil {
		return fmt.Errorf("issue parsing JSON from response: %v", err)
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", data.Name)

	caught := rand.Intn(2)
	if caught == 0 {
		fmt.Printf("%s escaped!\n", data.Name)
		return nil
	}

	c.Pokedex[data.Name] = data
	fmt.Printf("%s was caught!\nYou may now inspect it with the inspect command.\n", data.Name)

	return nil
}

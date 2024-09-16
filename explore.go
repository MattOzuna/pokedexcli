package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func explore(c *config, input string) error {
	reqURL := "https://pokeapi.co/api/v2/location-area/" + input

	cacheData, inCache := c.Cache.Get(reqURL)
	var data LocationRes
	if !inCache {
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

		c.Cache.Add(reqURL, resData)

	} else {
		if err := json.Unmarshal(cacheData, &data); err != nil {
			return fmt.Errorf("issue parsing JSON from cache: %v", err)
		}
	}
	fmt.Printf("Exploring %v...\nFound Pokemon:\n", data.Name)

	for _, val := range data.PokemonEncounters {
		fmt.Println(" -", val.Pokemon.Name)
	}

	return nil
}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func mapb(c *config) error {
	if c.Previous == "" {
		return fmt.Errorf("unable to get previous locations, type map to get next 20 locations")
	}
	reqURL := c.Previous

	res, err := http.Get(reqURL)
	if err != nil {
		return fmt.Errorf("issue reaching API: %v", err)
	}
	defer res.Body.Close()

	var data PokedexRes
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&data); err != nil {
		return fmt.Errorf("issue parsing JSON: %v", err)
	}

	c.Next = data.Next
	c.Previous = data.Previous

	for _, locations := range data.Results {
		fmt.Println(locations.Name)
	}

	return nil
}

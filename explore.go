package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func explore(c *config, input string) error {
	// fmt.Println(c.Next)

	reqURL := "https://pokeapi.co/api/v2/location-area/" + input
	fmt.Println(reqURL)

	// cacheData, inCache := c.Cache.Get(input)
	var data map[string]interface{}
	// if !inCache {
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
	fmt.Println(data)
	// c.Cache.Add(reqURL, resData)

	// } else {
	// 	if err := json.Unmarshal(cacheData, &data); err != nil {
	// 		return fmt.Errorf("issue parsing JSON from cache: %v", err)
	// 	}
	// }

	return nil
}

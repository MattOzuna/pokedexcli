package main

import (
	"time"

	"github.com/MattOzuna/pokedexcli/internal/pokecache"
)

type config struct {
	Next     string
	Previous string
	Cache    pokecache.Cache
	Pokedex  Pokedex
}

func intialize() config {
	interval := 10 * time.Second
	cache := pokecache.NewCache(interval)
	return config{
		Next:     "",
		Previous: "",
		Cache:    cache,
		Pokedex: Pokedex{
			Pokemon: []Pokemon{},
		},
	}
}

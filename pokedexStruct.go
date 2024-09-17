package main

type PokedexRes struct {
	Count    int                `json:"count"`
	Next     string             `json:"next"`
	Previous string             `json:"previous"`
	Results  []NamedAPIResource `json:"results"`
}

type NamedAPIResource struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type LocationRes struct {
	Name                 string                   `json:"name"`
	Id                   int                      `json:"id"`
	GameIndex            int                      `json:"game_index"`
	EncounterMethodRates []map[string]interface{} `json:"encounter_method_rates"`
	Location             map[string]interface{}   `json:"location"`
	Names                []map[string]interface{} `json:"names"`
	PokemonEncounters    []PokemonEncounter       `json:"pokemon_encounters"`
}

type PokemonEncounter struct {
	Pokemon        NamedAPIResource         `json:"pokemon"`
	VersionDetails []map[string]interface{} `json:"version_details"`
}

type Pokemon struct {
	Name                   string                   `json:"name"`
	LocationAreaEncounters string                   `json:"location_area_encounters"`
	Id                     int                      `json:"id"`
	BaseExperience         int                      `json:"base_experience"`
	Height                 int                      `json:"height"`
	IsDefault              bool                     `json:"is_default"`
	Order                  int                      `json:"order"`
	Weight                 int                      `json:"weight"`
	Abilities              []map[string]interface{} `json:"abilities"`
	Forms                  []map[string]interface{} `json:"forms"`
	HeldItems              []map[string]interface{} `json:"held_items"`
	Moves                  []map[string]interface{} `json:"moves"`
	Species                NamedAPIResource         `json:"species"`
	Stats                  []map[string]interface{} `json:"stats"`
	Types                  []map[string]interface{} `json:"types"`
}

type Pokedex struct {
	Pokemon []Pokemon
}

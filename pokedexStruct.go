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

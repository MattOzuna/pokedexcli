package main

type PokedexRes struct {
	Count    int         `json:"count"`
	Next     string      `json:"next"`
	Previous string      `json:"previous"`
	Results  []Locations `json:"results"`
}

type Locations struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type LocationRes struct {
	Name                 string              `json:"name"`
	Id                   int                 `json:"id"`
	GameIndex            int                 `json:"game_index"`
	EncounterMethodRates EncounterMethodRate `json:"encounter_method_rates"`
}

type EncounterMethodRate struct {
}

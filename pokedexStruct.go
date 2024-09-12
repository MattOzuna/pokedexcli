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

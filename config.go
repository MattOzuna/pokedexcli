package main

type config struct {
	Next     string
	Previous string
}

func intialize() config {
	return config{
		Next:     "",
		Previous: "",
	}
}

package main

import (
	"pokedex/internal/pokeapi"
	"time"
)

func main() {
	pokeapiClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	c := &config{
		pokeapiClient: pokeapiClient,
		pokedex:       make(map[string]Pokemon),
	}

	startRepl(c)
}

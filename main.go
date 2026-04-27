package main

import (
	"pokedex/internal/pokeapi"
	"time"
)

func main() {
	pokeapiClient := pokeapi.NewClient(5*time.Second, time.Minute*5)
	c := &config{
		pokeapiClient: pokeapiClient,
	}

	startRepl(c)
}

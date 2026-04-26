package main

import (
	"pokedex/internal/pokeapi"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 10*time.Second)
	c := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(c)
}

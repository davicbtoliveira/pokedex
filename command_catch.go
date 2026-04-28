package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, param *string) error {
	if param == nil {
		return fmt.Errorf("you must provide a pokemon name")
	}

	pokeInfo, err := cfg.pokeapiClient.PokemonInfo(param)
	if err != nil {
		return err
	}

	baseXP := pokeInfo.BaseExperience

	fmt.Printf("Throwing a Pokeball at %v...\n", pokeInfo.Name)

	if rand.Intn(350) > baseXP {
		caught := Pokemon{
			Name:           pokeInfo.Name,
			BaseExperience: pokeInfo.BaseExperience,
			Height:         pokeInfo.Height,
			IsDefault:      pokeInfo.IsDefault,
			Order:          pokeInfo.Order,
			Weight:         pokeInfo.Weight,
		}
		cfg.pokedex["caughts"] = append(cfg.pokedex["caughts"], caught)
		fmt.Printf("%v was caught!\n", caught.Name)
		fmt.Printf("Now you can see %v in your pokedex, by using the pokedex command.\n", caught.Name)
		return nil
	}

	return fmt.Errorf("%v escaped!", pokeInfo.Name)
}

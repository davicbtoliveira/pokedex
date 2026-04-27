package main

import "fmt"

func commandInspect(cfg *config, param *string) error {
	pokemon, err := cfg.pokeapiClient.PokemonInfo(param)
	if err != nil {
		return err
	}

	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, v := range pokemon.Stats {
		fmt.Printf(" -%s: %v\n", v.Stat.Name, v.BaseStat)
	}
	fmt.Println("Types:")
	for _, v := range pokemon.Types {
		fmt.Printf(" - %v\n", v.Type.Name)
	}

	return nil
}

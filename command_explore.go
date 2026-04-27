package main

import "fmt"

func commandExplore(cfg *config, area *string) error {
	if area == nil {
		return fmt.Errorf("you need to inform a location area")
	}

	locationInfo, err := cfg.pokeapiClient.LocationInfo(area)
	if err != nil {
		return err
	}

	for _, v := range locationInfo.PokemonEncounters {
		fmt.Println(v.Pokemon.Name)
	}

	return nil
}

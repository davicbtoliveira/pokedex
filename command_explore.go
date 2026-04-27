package main

import "fmt"

func commandExplore(cfg *config, area *string) error {
	if area == nil {
		return fmt.Errorf("you must provide a location name")
	}

	locationInfo, err := cfg.pokeapiClient.LocationInfo(area)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", locationInfo.Name)
	fmt.Println("Found Pokemon: ")
	for _, enc := range locationInfo.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}

	return nil
}

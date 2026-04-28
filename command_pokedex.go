package main

import "fmt"

func commandPokedex(cfg *config, param *string) error {
	fmt.Println("Your Pokedex:")
	if len(cfg.pokedex) == 0 {
		return fmt.Errorf("you don't have any pokemon")
	}
	for _, v := range cfg.pokedex["caughts"] {
		fmt.Printf(" - %v\n", v.Name)
	}

	return nil
}

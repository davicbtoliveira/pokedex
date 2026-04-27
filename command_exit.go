package main

import (
	"fmt"
	"os"
)

func commandExit(c *config, param *string) error {
	fmt.Printf("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

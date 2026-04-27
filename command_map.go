package main

import (
	"errors"
	"fmt"
)

func commandMap(c *config, param *string) error {
	locationsResp, err := c.pokeapiClient.ListLocations(c.Next)
	if err != nil {
		return err
	}

	c.Next = locationsResp.Next
	c.Previous = locationsResp.Previous

	for _, v := range locationsResp.Results {
		fmt.Println(v.Name)
	}

	return nil
}

func commandMapb(c *config, param *string) error {
	if c.Previous == nil {
		return errors.New("you're on the first page")
	}

	locationsResp, err := c.pokeapiClient.ListLocations(c.Previous)
	if err != nil {
		return nil
	}

	c.Next = locationsResp.Next
	c.Previous = locationsResp.Previous

	for _, v := range locationsResp.Results {
		fmt.Println(v.Name)
	}

	return nil
}

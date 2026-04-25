package main

import (
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

var commands = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the pokedex",
		callback:    commandExit,
	},
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	result := strings.Fields(text)

	return result
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!\nUsage:\n")
	for k, v := range commands {
		fmt.Printf("%v: %v\n", k, v.description)
	}
	return nil
}

func commandExit() error {
	fmt.Printf("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
}

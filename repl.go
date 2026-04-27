package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedex/internal/pokeapi"
	"strings"
)

type Pokemon struct {
	Name           string
	BaseExperience int
	Height         int
	IsDefault      bool
	Order          int
	Weight         int
}

type config struct {
	pokeapiClient pokeapi.Client
	pokedex       map[string]Pokemon
	Next          *string
	Previous      *string
}

func startRepl(c *config) {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		var param *string
		if len(words) > 1 {
			param = &words[1]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(c, param)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	result := strings.Fields(text)
	return result
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, *string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the Previous page of locatons",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "List all pokemons in a given area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a given pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspect a given Pokemon",
			callback:    commandInspect,
		},
	}
}

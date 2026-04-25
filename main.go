package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		rawInput := scanner.Text()
		clearedInput := cleanInput(rawInput)

		commands = getCommands()
		command, ok := commands[clearedInput[0]]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		err := command.callback()
		if err != nil {
			fmt.Println(err)
		}
	}
}

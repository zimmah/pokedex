package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/zimmah/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient		pokeapi.Client
	nextLocationsURL	*string
	prevLocationsURL	*string
}

func repl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 { continue }

		commandName := words[0]

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg)
			if err != nil { fmt.Println(err) }
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct {
	name			string
	description		string
	callback func(*config)	error
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
			name: 		 "map",
			description: "Get next page of locations",
			callback:	 commandMap,
		},
		"mapb": {
			name: 		 "mapb",
			description: "Get prev page of locations",
			callback:	 commandMapb,
		},
	}
}

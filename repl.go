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
	caughtPokemon		map[string]pokeapi.Pokemon
}

func repl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 { continue }

		commandName := words[0]
		var parameter string

		if len(words) > 1 {
			parameter = words[1]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, parameter)
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
	callback func(*config, string)	error
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
		"explore": {
			name: 		"explore",
			description: "Find pokemon at location",
			callback: commandExplore,
		},
		"catch": {
			name: "catch",
			description: "Throw pokeball to attempt a catch of pokemon",
			callback: commandCatch,
		},
		"inspect": {
			name: "inspect",
			description: "Get details about a caught pokemon",
			callback: commandInspect,
		},
	}
}

package main

import (
	"fmt"
	"time"

	"github.com/zimmah/pokedex/internal/pokeapi"
)

func main() {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("type help for help, exit to exit")
	pokeClient := pokeapi.NewClient(5 * time.Second, 5 * time.Minute)
	cfg := &config{
		caughtPokemon: map[string]pokeapi.Pokemon{},
		pokeapiClient: pokeClient,
	}
	
	repl(cfg)
}
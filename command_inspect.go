package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, name string) error {
	pokemon, isCaught := cfg.caughtPokemon[name]
	if !isCaught { return errors.New("you must catch this Pokemon first")}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n",pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Stats: \n")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Printf("Types: \n")
	for _, pokemonType := range pokemon.Types {
		fmt.Printf("  - %s\n", pokemonType.Type.Name)
	}

	return nil
}
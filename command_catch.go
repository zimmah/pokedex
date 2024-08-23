package main

import (
	"fmt"
	"errors"
	"math/rand"
)

func commandCatch(cfg *config, name string) error {
	if name == "" {
		return errors.New("you must provide a pokemon name")
	}

	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil { return err }

	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	if pokemon.BaseExperience <= rand.Intn(1000) {
		fmt.Printf("%s was caught!\n", name)
	} else {
		fmt.Printf("%s escaped!\n", name)
	}

	return nil
}
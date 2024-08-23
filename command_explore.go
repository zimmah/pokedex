package main

import (
	"fmt"
	"errors"
)

func commandExplore(cfg *config, locationArg string) error {
	if locationArg == "" {
		return errors.New("you must provide a location name")
	}

	name := locationArg
	location, err := cfg.pokeapiClient.GetLocation(name)
	if err != nil { return err }

	fmt.Printf("Exploring %s...\n", location.Name)
	fmt.Println("Found Pokemon: ")
	for _, enc := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", enc.Pokemon.Name)
	}
	return nil
}
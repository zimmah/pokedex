package main

import (
	"fmt"
	"errors"
)

func commandMap(cfg *config, _ string) error {
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil { return err }

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	for _, location := range locationsResp.Result {
		fmt.Printf("Location: %s: (%s)\n", location.Name, location.Url)
	}

	return nil
}

func commandMapb(cfg *config, _ string) error {
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil { return err }

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	for _, location := range locationsResp.Result {
		fmt.Printf("Location: %s: (%s)\n", location.Name, location.Url)
	}

	return nil
}
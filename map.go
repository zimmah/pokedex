package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"strconv"
)

const apiRoot string = "https://pokeapi.co/api/v2/location-area/"
var offset = 0

type RespLocations struct {
	Count int `json:"count"`
	Next *string `json:"next"`
	Previous *string `json:"previous"`
	Result []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

func commandMap(cfg *config) error {
	res, err := http.Get(apiRoot + "?offset=" + strconv.Itoa(offset) + "&limit=20")
	if err != nil { return fmt.Errorf("error creating request: %w", err) }

	defer res.Body.Close()

	if res.StatusCode > 299 {return fmt.Errorf("non-OK status code: %d", res.StatusCode)}

	var locations RespLocations
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&locations)
	if err != nil {
		fmt.Println(res.Body)
		return fmt.Errorf("error parsing data: %w", err)
	}

	for _, location := range locations.Result {
		fmt.Printf("Location: %s: (%s)\n", location.Name, location.Url)
	}

	return nil
}

func commandMapb(cfg *config) error {
	return nil
}
package main

import(
	"fmt"
	"os"
)

func commandExit(cfg *config, _ string) error{
	fmt.Print("Thank you for using Pokedex!")
	os.Exit(0)
	return nil
}
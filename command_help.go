package main

import("fmt")

func commandHelp(cfg *config) error {
	for _, cliCommand := range getCommands() {
		fmt.Printf("%s: %s\n", cliCommand.name, cliCommand.description)
	}
	fmt.Println()
	return nil
}
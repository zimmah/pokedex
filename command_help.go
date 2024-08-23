package main

import("fmt")

func commandHelp(cfg *config, _ string) error {
	for _, cliCommand := range getCommands() {
		fmt.Printf("%s: %s\n", cliCommand.name, cliCommand.description)
	}
	fmt.Println()
	return nil
}
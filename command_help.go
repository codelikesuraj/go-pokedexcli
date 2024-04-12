package main

import "fmt"

func commandHelp(cfg *config, _ ...string) error {
	fmt.Println()

	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")

	for _, cmd := range getCommands() {
		fmt.Printf("\t%s: %s\n", cmd.name, cmd.description)
	}

	fmt.Println()

	return nil
}

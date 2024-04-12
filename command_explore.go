package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("a location is required")
	}

	resp, err := cfg.pokeapiClient.GetLocation(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", resp.Location.Name)

	if l := len(resp.PokemonEncounters); l > 0 {
		fmt.Println("Found", l, "pokemon(s):")
		for _, enc := range resp.PokemonEncounters {
			fmt.Printf(" - %s\n", enc.Pokemon.Name)
		}
	} else {
		fmt.Println("No pokemons found!")
	}

	fmt.Println()
	return nil
}

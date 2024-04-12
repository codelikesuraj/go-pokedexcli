package main

import (
	"fmt"
)

func commandPokedex(cfg *config, _ ...string) error {
	if len(cfg.caughtPokemon) < 1 {
		fmt.Println("you have not caught any pokemon!")
	} else {
		fmt.Println("Your pokedex:")
		for name := range cfg.caughtPokemon {
			fmt.Println(" -", name)
		}
	}

	fmt.Println()
	return nil
}

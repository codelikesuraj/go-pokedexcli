package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("a pokemon is required")
	}

	pokemon, ok := cfg.caughtPokemon[args[0]]
	if !ok {
		fmt.Println("you have not caught that pokemon!")
		fmt.Println()
		return nil
	}

	fmt.Println("Name:", pokemon.Name)
	fmt.Println("Height:", pokemon.Height)
	fmt.Println("Weight:", pokemon.Weight)
	fmt.Println("Stats:")
	for _, s := range pokemon.Stats {
		fmt.Printf(" -%s: %d\n", s.Stat.Name, s.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Println(" -", t.Type.Name)
	}

	fmt.Println()
	return nil
}

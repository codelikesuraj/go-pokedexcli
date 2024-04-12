package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("a pokemon is required")
	}

	pokemon, err := cfg.pokeapiClient.GetPokemon(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a ball at %s...\n", pokemon.Name)
	if rand.Intn(pokemon.BaseExperience) < 40 {
		fmt.Println(pokemon.Name, "was caught!")
		cfg.caughtPokemon[pokemon.Name] = pokemon
	} else {
		fmt.Println(pokemon.Name, "could not be caught!")
	}

	fmt.Println()

	return nil
}

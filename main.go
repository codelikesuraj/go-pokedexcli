package main

import (
	"time"

	"github.com/codelikesuraj/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAreaUrl *string
	prevLocationAreaUrl *string
	caughtPokemon       map[string]pokeapi.Pokemon
}

func main() {
	cfg := config{
		caughtPokemon: map[string]pokeapi.Pokemon{},
		pokeapiClient: pokeapi.NewClient(time.Minute),
	}

	startRepl(&cfg)
}

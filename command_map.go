package main

import (
	"errors"
	"fmt"

	"github.com/codelikesuraj/pokedexcli/internal/pokeapi"
)

func commandMap(cfg *config) error {
	resp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationAreaUrl)
	if err != nil {
		return err
	}

	cfg.nextLocationAreaUrl = resp.Next
	cfg.prevLocationAreaUrl = resp.Previous

	displayLocations(resp.Locations)

	fmt.Println()
	return nil
}

func commandMapB(cfg *config) error {
	if cfg.prevLocationAreaUrl == nil {
		baseUrl := fmt.Sprintf(
			"%s%s",
			*pokeapi.GetBaseURL(),
			"/location",
		)
		cfg.nextLocationAreaUrl = &baseUrl
		return errors.New("you are on the first page")
	}

	resp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationAreaUrl)
	if err != nil {
		return err
	}

	cfg.nextLocationAreaUrl = resp.Next
	cfg.prevLocationAreaUrl = resp.Previous

	displayLocations(resp.Locations)
	fmt.Println()

	return nil
}

func displayLocations(locations []pokeapi.LocationArea) {
	fmt.Println("Location areas:")
	for _, loc := range locations {
		fmt.Printf(" - %s\n", loc.URL)
	}
}

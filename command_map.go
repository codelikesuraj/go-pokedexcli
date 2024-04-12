package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/codelikesuraj/pokedexcli/internal/pokeapi"
)

func commandMap(cfg *config, _ ...string) error {
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

func commandMapB(cfg *config, _ ...string) error {
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
		fmt.Printf(" - %s (%s)\n", loc.Name, getId(loc.URL))
	}
}

func getId(url string) string {
	return strings.Trim(strings.TrimPrefix(url, *pokeapi.GetBaseURL()+"/location/"), "/")
}

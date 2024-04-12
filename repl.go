package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "displays the help menu",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "exit the pokedexcli",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "list location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "list previous location areas",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "explore a location",
			callback:    commandExplore,
		},
	}
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		cleaned := cleanInput(scanner.Text())

		if len(cleaned) == 0 {
			continue
		}

		command, ok := getCommands()[cleaned[0]]
		if !ok {
			displayError("Invalid command!", true)
			continue
		}

		start := time.Now()
		if err := command.callback(cfg, cleaned[1:]...); err != nil {
			displayError(err.Error(), false)
		}
		fmt.Println("Interval:", time.Since(start).Seconds(), "seconds")
	}
}

func cleanInput(s string) []string {
	return strings.Fields(strings.ToLower(s))
}

func displayError(message string, showHint bool) {
	fmt.Println()
	fmt.Println(message)
	if showHint {
		fmt.Println("Type 'help' to show available commands.")
	}
	fmt.Println()
}

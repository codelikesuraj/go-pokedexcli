package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
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
	}
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		cleaned := cleanInput(scanner.Text())

		if len(cleaned) == 0 {
			continue
		}

		command, ok := getCommands()[cleaned[0]]
		if ok {
			command.callback()
			continue
		}

		invalidCommand()
	}
}

func cleanInput(s string) []string {
	return strings.Fields(strings.ToLower(s))
}

func invalidCommand() {
	fmt.Println()
	fmt.Println("Invalid command!")
	fmt.Println("Type 'help' to show available commands.")
	fmt.Println()
}

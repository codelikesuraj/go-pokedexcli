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
		print("pokedexcli> ")
		scanner.Scan()
		text := scanner.Text()

		if len(text) == 0 {
			continue
		}

		command, ok := getCommands()[cleanInput(text)[0]]
		if ok {
			command.callback()
			continue
		}

		fmt.Println("Invalid command. Type 'help' to show available commands!")
	}
}

func cleanInput(s string) []string {
	return strings.Fields(strings.ToLower(s))

}

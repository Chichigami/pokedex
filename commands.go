package main

import (
	"fmt"
	"os"
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
			description: "Display help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the pokedex",
			callback:    commandExit,
		},
	}
}

func commandHelp() error {
	commands := "Usage: \n"
	for k := range getCommands() {
		commands += fmt.Sprintf("%s: %s\n", getCommands()[k].name, getCommands()[k].description)
	}
	fmt.Println(commands)
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}

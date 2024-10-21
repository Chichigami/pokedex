package main

import pokecache "github.com/chichigami/pokedex/internal"

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
		"map": {
			name:        "map",
			description: "Displays 20 locations. Can use again to get the next 20.",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Opposite of map command. Will display previous 20 locations.",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Explore the given region",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch that pokemon",
			callback:    commandCatch,
		},
	}
}

type Config struct {
	Next           string
	Previous       string
	cache          *pokecache.Cache
	caughtPokemons map[string]Pokemon
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, ...string) error
}

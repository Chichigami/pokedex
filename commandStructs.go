package main

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
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
	}
}

type Config struct {
	Next     string
	Previous string
}

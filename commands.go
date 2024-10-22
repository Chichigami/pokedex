package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func commandHelp(cfg *Config, args ...string) error {
	commands := "Usage: \n"
	for _, command := range getCommands() {
		commands += fmt.Sprintf("%s: %s\n", command.name, command.description)
	}
	fmt.Println(commands)
	return nil
}

func commandExit(cfg *Config, args ...string) error {
	os.Exit(0)
	return nil
}

func commandMap(cfg *Config, args ...string) error {
	if len(args) != 0 {
		return fmt.Errorf("can't use any arguments with this command")
	}
	if cfg.Next == "" {
		cfg.Next = "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"
	}

	var currentLocations area
	var body []byte
	body, ok := cfg.cache.Get(cfg.Next)
	if !ok {
		res, res_err := http.Get(cfg.Next)
		if res_err != nil {
			fmt.Println(res_err)
		}
		body, _ = io.ReadAll(res.Body)
		defer res.Body.Close()
		cfg.cache.Add(cfg.Next, body)
	}

	err := json.Unmarshal(body, &currentLocations)
	if err != nil {
		fmt.Println(err)
	}
	cfg.Next = currentLocations.Next
	cfg.Previous = currentLocations.Previous
	var twentyLocations string
	for _, location := range currentLocations.Results {
		twentyLocations += location.Name + "\n"
	}
	fmt.Println(twentyLocations)
	return nil
}

func commandMapb(cfg *Config, args ...string) error {
	if len(args) != 0 {
		return fmt.Errorf("can't use any arguments with this command")
	}
	if cfg.Previous == "" {
		return fmt.Errorf("can't go backwards")
	}
	var currentLocations area
	var body []byte
	var twentyLocations string

	body, ok := cfg.cache.Get(cfg.Previous)
	if !ok {
		res, res_err := http.Get(cfg.Previous)
		if res_err != nil {
			fmt.Println(res_err)
		}
		body, _ = io.ReadAll(res.Body)
		defer res.Body.Close()
		cfg.cache.Add(cfg.Previous, body)
	}
	cfg.Next = currentLocations.Next
	cfg.Previous = currentLocations.Previous

	err := json.Unmarshal(body, &currentLocations)
	if err != nil {
		fmt.Println(err)
	}
	for _, location := range currentLocations.Results {
		twentyLocations += location.Name + "\n"
	}
	fmt.Println(twentyLocations)
	return nil
}

func commandExplore(cfg *Config, args ...string) error {
	switch len(args) {
	case 0:
		return fmt.Errorf("can't use explore without a location")
	case 1:
		var pokemons string
		var area area_info
		var body []byte
		location := args[0]
		fmt.Printf("Exploring %s... \n", location)
		url := "https://pokeapi.co/api/v2/location-area/" + location

		body, ok := cfg.cache.Get(url)
		if !ok {
			res, err := http.Get(url)
			if err != nil {
				fmt.Println(err)
			}
			body, _ = io.ReadAll(res.Body)
			defer res.Body.Close()
			cfg.cache.Add(url, body)
		}
		fmt.Println("Found Pokemon(s):")
		err := json.Unmarshal(body, &area)
		if err != nil {
			fmt.Println(err)
		}

		for _, pokemonStruct := range area.PokemonEncounters {
			pokemons += pokemonStruct.Pokemon.Name + "\n"
		}
		fmt.Println(pokemons)
		return nil
	default:
		return fmt.Errorf("can't explore multiple location at the same time")
	}
}

func commandCatch(cfg *Config, args ...string) error {
	var pokemon Pokemon
	target := strings.ToLower(args[0])
	url := "https://pokeapi.co/api/v2/pokemon/" + target
	_, ok := cfg.caughtPokemons[target]
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
		}
		body, _ := io.ReadAll(res.Body)
		defer res.Body.Close()
		err = json.Unmarshal(body, &pokemon)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Throwing a pokeball at %s...\n", target)
		if pokemon.Caught() {
			fmt.Printf("%s was caught! \n", pokemon.Name)
			cfg.caughtPokemons[pokemon.Name] = pokemon
		} else {
			fmt.Printf("%s escaped!\n", pokemon.Name)
		}
	} else {
		fmt.Println("pokemon already caught")
	}
	return nil
}

func commandInspect(cfg *Config, args ...string) error {
	pokemon, ok := cfg.caughtPokemons[args[0]]
	if !ok {
		fmt.Println("you have not caught that pokemon")
	} else {
		var typesUnorderedList string
		var statUnorderedList string
		for _, typeStruct := range pokemon.Types {
			typesUnorderedList += fmt.Sprintf("  - %s\n", typeStruct.Type.Name)
		}
		for _, statStruct := range pokemon.Stats {
			statUnorderedList += fmt.Sprintf("  -%s: %d\n", statStruct.Stat.Name, statStruct.BaseStat)
		}
		fmt.Printf("Name: %s\n"+
			"Height: %v\n"+
			"Weight: %v\n"+
			"Stats:\n"+
			"%s"+
			"Types:\n"+
			"%s\n", pokemon.Name, pokemon.Height, pokemon.Weight, statUnorderedList, typesUnorderedList)
	}
	return nil
}

func commandPokedex(cfg *Config, args ...string) error {
	if len(args) != 0 {
		return fmt.Errorf("can't use any arguments with this command")
	}
	var pokemonUnorderedList string
	for _, pokemon := range cfg.caughtPokemons {
		pokemonUnorderedList += "  - " + pokemon.Name + "\n"
	}
	fmt.Printf("Your Pokedex:\n%s", pokemonUnorderedList)
	return nil
}

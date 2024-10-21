package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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
		return fmt.Errorf("can't use map with any arguments")
	}
	if cfg.Next == "" {
		cfg.Next = "https://pokeapi.co/api/v2/location-area/"
	}
	var currentLocations area
	res, res_err := http.Get(cfg.Next)
	if res_err != nil {
		fmt.Println(res_err)
	}
	body, _ := io.ReadAll(res.Body)
	defer res.Body.Close()
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
		return fmt.Errorf("can't use mapb with any arguments")
	}
	if cfg.Previous == "" {
		return fmt.Errorf("can't go backwards")
	}
	var currentLocations area
	res, res_err := http.Get(cfg.Previous)
	if res_err != nil {
		fmt.Println(res_err)
	}
	body, _ := io.ReadAll(res.Body)
	defer res.Body.Close()
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

func explore(cfg *Config, args ...string) error {
	switch len(args) {
	case 0:
		return fmt.Errorf("can't use explore without a location")
	case 1:
		var pokemons string
		var area area_info
		var body []byte
		location := args[0]
		fmt.Printf("Exploring %s... \n", location)
		url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", location)

		body, ok := cfg.cache.Get(url)
		if !ok {
			fmt.Println("not cached")
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

		for _, pokemon_struct := range area.PokemonEncounters {
			pokemons += pokemon_struct.Pokemon.Name + "\n"
		}
		fmt.Println(pokemons)
		return nil
	default:
		return fmt.Errorf("can't explore multiple location at the same time")
	}
}

// func getReqParser(url string, s struct{}) struct{} {
// 	return struct{}
// }

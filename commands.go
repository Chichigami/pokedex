package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func commandHelp(cfg *Config) error {
	commands := "Usage: \n"
	for k := range getCommands() {
		commands += fmt.Sprintf("%s: %s\n", getCommands()[k].name, getCommands()[k].description)
	}
	fmt.Println(commands)
	return nil
}

func commandExit(cfg *Config) error {
	os.Exit(0)
	return nil
}

func commandMap(cfg *Config) error {
	if cfg.Next == "" {
		cfg.Next = "https://pokeapi.co/api/v2/location/"
	}
	var currentLocations locations
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

func commandMapb(cfg *Config) error {
	if cfg.Previous == "" {
		return errors.New("can't go backwards")
	}
	var currentLocations locations
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

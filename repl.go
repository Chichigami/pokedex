package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	pokecache "github.com/chichigami/pokedex/internal"
)

func repl() {
	fmt.Println("Welcome to the Pokedex!")
	var config = Config{
		cache:          pokecache.NewCache(10 * time.Minute),
		caughtPokemons: map[string]Pokemon{},
	}

	for {
		var input string
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			input = scanner.Text()
		}
		splitted_input := cleanInput(input)
		command := splitted_input[0]
		if _, ok := getCommands()[command]; !ok {
			fmt.Println("Unknown command")
		} else {
			var err error
			if len(splitted_input) > 1 {
				err = getCommands()[splitted_input[0]].callback(&config, splitted_input[1])
			} else {
				err = getCommands()[splitted_input[0]].callback(&config)
			}
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

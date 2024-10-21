package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	pokecache "github.com/chichigami/pokedex/internal"
)

func main() {
	fmt.Println("Welcome to the Pokedex")
	var config = Config{
		cache:          pokecache.NewCache(10 * time.Minute),
		caughtPokemons: map[string]Pokemon{},
	}

	for {
		//take an input
		var input string
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Pokedex > ")
		if scanner.Scan() {
			input = scanner.Text()
		}
		input = strings.TrimSpace(strings.ToLower(input))
		splitted_input := strings.SplitN(input, " ", 2)
		//use input
		if _, ok := getCommands()[splitted_input[0]]; !ok {
			fmt.Println("command not found")
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

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Welcome to the Pokedex")
	commands := getCommands()
	for {
		var input string
		fmt.Print("Pokedex > ")
		fmt.Scanln(&input)
		input = strings.ToLower(input)
		if _, ok := commands[input]; !ok {
			fmt.Println("command not found")
		} else {
			commands[input].callback()
		}
	}
}

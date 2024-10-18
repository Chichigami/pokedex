package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Welcome to the Pokedex")
	var config = Config{}
	for {	
		var input string
		fmt.Print("Pokedex > ")
		fmt.Scanln(&input)
		input = strings.ToLower(input)
		if _, ok := getCommands()[input]; !ok {
			fmt.Println("command not found")
		} else {
			command_err := getCommands()[input].callback(&config)
			if command_err != nil {
				fmt.Println(command_err)
			}
		}
	}
}

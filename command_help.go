package main

import "fmt"

func help() {
	fmt.Println("--------------------------------------------------")
	fmt.Println("Welcome to the pokedexCLI!")
	fmt.Println("Here are the available commands:")
	commands := getCommands()
	for _, cmd := range commands {
		fmt.Printf("%s : %s\n", cmd.name, cmd.description)
	}
	fmt.Println("--------------------------------------------------")
}

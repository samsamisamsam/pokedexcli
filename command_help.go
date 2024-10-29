package main

import "fmt"

func help() error {
	availableCommands := getCommands()
	fmt.Println("------------------------------------")
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Here are the available commands:")
	for _, cmd := range availableCommands {
		fmt.Printf("-- %s: %s\n", cmd.name, cmd.description)

	}
	fmt.Println("------------------------------------")
	return nil
}

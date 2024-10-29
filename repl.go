package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Pokedex>")
		scanner.Scan()

		text := scanner.Text()
		cleaned := cleanInput(text)
		commandName := cleaned[0]

		availableCommands := getCommands()
		command, ok := availableCommands[commandName]
		if !ok {
			fmt.Println("invalid command")
			continue
		}
		command.callback()
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}

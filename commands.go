package main

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    exitCallback,
		},
		"help": {
			name:        "help",
			description: "Displays available commands",
			callback:    helpCallback,
		},
		"map": {
			name:        "map",
			description: " Displays a list of 20 locations",
			callback:    mapCallback,
		},
	}
}

package main

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "  exit",
			description: "Exit the Pokedex",
			callback:    exitCallback,
		},
		"help": {
			name:        "  help",
			description: "Displays available commands",
			callback:    helpCallback,
		},
		"map": {
			name:        "   map",
			description: " Displays a list of 20 locations",
			callback:    mapCallback,
		},
		"mapb": {
			name:        "  mapb",
			description: "Displays a list of the 20 previous locations",
			callback:    mapbCallback,
		},
		"explore": {
			name:        "explore {location_area}",
			description: "Displays wild pokemon who might be appearing in a location area",
			callback:    exploreCallback,
		},
		"catch": {
			name:        " catch {pokemon_name}",
			description: "Try to catch a pokemon",
			callback:    catchCallback,
		},
	}
}

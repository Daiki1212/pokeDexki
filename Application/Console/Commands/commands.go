package Commands

import "github.com/Daiki1212/pokeDexki/Services/Pokeapi"

type cliCommand struct {
	name        string
	description string
	Callback    func(config *Pokeapi.Config, queryParam string, pokedex *map[string]Pokeapi.Pokemon) error
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			Callback:    CommandHelp,
		},
		"map": {
			name:        "map",
			description: "Get the next page of locations",
			Callback:    CommandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Get the previous page of locations",
			Callback:    CommandMapb,
		},
		"explore": {
			name:        "explore <location>",
			description: "Explore the Pokemons of a location! <location> needs to be the id or name of a location",
			Callback:    CommandExplore,
		},
		"catch": {
			name:        "catch <pokemon>",
			description: "Tries to Catch a Pokemon! <pokemon> needs to be the name of a Pokemon",
			Callback:    CommandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon>",
			description: "Inspects a Pokemon! <pokemon> needs to be the name of a Pokemon",
			Callback:    CommandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Displays all Pokemon in your Pokedex",
			Callback:    CommandPokedex,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			Callback:    CommandExit,
		},
	}
}

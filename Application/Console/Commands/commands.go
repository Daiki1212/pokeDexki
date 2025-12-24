package Commands

import "github.com/Daiki1212/pokeDexki/Services/Pokeapi"

type cliCommand struct {
	name        string
	description string
	Callback    func(config *Pokeapi.Config) error
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
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			Callback:    CommandExit,
		},
	}
}

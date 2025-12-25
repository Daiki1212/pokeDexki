package Commands

import (
	"fmt"

	"github.com/Daiki1212/pokeDexki/Services/Pokeapi"
)

func CommandHelp(_ *Pokeapi.Config, _ string, _ *map[string]Pokeapi.Pokemon) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, cmd := range GetCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}

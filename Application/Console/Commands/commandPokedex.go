package Commands

import (
	"fmt"

	"github.com/Daiki1212/pokeDexki/Services/Pokeapi"
)

func CommandPokedex(_ *Pokeapi.Config, _ string, pokedex *map[string]Pokeapi.Pokemon) error {
	fmt.Println("Your Pokedex:")
	for pokemonName := range *pokedex {
		fmt.Printf("\t- %s\n", pokemonName)
	}

	return nil
}

package Commands

import (
	"fmt"

	"github.com/Daiki1212/pokeDexki/Services/Pokeapi"
)

func CommandInspect(_ *Pokeapi.Config, pokemonName string, pokedex *map[string]Pokeapi.Pokemon) error {
	pokemon, exists := (*pokedex)[pokemonName]
	if !exists {
		return fmt.Errorf("pokemon %s not found in Pokedex", pokemonName)
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Base Experience: %d\n", pokemon.BaseExperience)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("\t- %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typ := range pokemon.Types {
		fmt.Printf("\t- %s\n", typ.Type.Name)
	}

	return nil
}

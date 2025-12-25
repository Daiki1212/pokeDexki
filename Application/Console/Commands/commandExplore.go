package Commands

import (
	"fmt"

	"github.com/Daiki1212/pokeDexki/Services/Pokeapi"
)

func CommandExplore(cfg *Pokeapi.Config, location string, _ *map[string]Pokeapi.Pokemon) error {
	locationList, err := cfg.PokeapiClient.GetLocation(location)
	if err != nil {
		return err
	}

	fmt.Printf("Looking for Pokemon in %s\n", locationList.Name)
	fmt.Printf("Found %d Pokemon:\n", len(locationList.PokemonEncounter))
	for _, item := range locationList.PokemonEncounter {
		fmt.Printf("\t- %s\n", item.Pokemon.Name)
	}

	return nil
}

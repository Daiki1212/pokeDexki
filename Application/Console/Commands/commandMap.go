package Commands

import (
	"errors"
	"fmt"

	"github.com/Daiki1212/pokeDexki/Services/Pokeapi"
)

func CommandMapf(cfg *Pokeapi.Config, _ string, _ *map[string]Pokeapi.Pokemon) error {
	locationResponse, err := cfg.PokeapiClient.ListLocations(cfg.NextLocationsURL)
	if err != nil {
		return err
	}

	updateUrls(cfg, locationResponse)

	for _, location := range locationResponse.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func CommandMapb(cfg *Pokeapi.Config, _ string, _ *map[string]Pokeapi.Pokemon) error {
	if cfg.PrevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationResponse, err := cfg.PokeapiClient.ListLocations(cfg.PrevLocationsURL)
	if err != nil {
		return err
	}

	updateUrls(cfg, locationResponse)

	for _, location := range locationResponse.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func updateUrls(cfg *Pokeapi.Config, locationResponse Pokeapi.ResponseLocations) {
	cfg.NextLocationsURL = locationResponse.Next
	cfg.PrevLocationsURL = locationResponse.Previous
}

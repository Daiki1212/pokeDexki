package Commands

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Daiki1212/pokeDexki/Services/Pokeapi"
)

func CommandCatch(cfg *Pokeapi.Config, pokemonName string, pokedex *map[string]Pokeapi.Pokemon) error {
	pokemon, err := cfg.PokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	catchTry := rand.Intn(pokemon.BaseExperience * 3)
	for i := 0; i <= 3; i++ {
		fmt.Printf(".")
		time.Sleep(time.Second * 1)
	}
	fmt.Printf("\n")

	if pokemon.BaseExperience > catchTry {
		fmt.Printf("%s escaped...\n", pokemon.Name)
		return nil
	}
	fmt.Printf("%s caught!\n", pokemon.Name)

	(*pokedex)[pokemon.Name] = pokemon

	return nil
}

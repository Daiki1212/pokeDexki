package main

import (
	"github.com/Daiki1212/pokeDexki/Services/Pokeapi"
)

func main() {
	cfg := Pokeapi.NewConfig()
	// TODO: put the pokedex in a specific Struct for better maintainability
	pokedex := map[string]Pokeapi.Pokemon{}
	startRepl(&cfg, &pokedex)
}

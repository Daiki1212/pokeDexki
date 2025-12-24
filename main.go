package main

import (
	"github.com/Daiki1212/pokeDexki/Services/Pokeapi"
)

func main() {
	cfg := Pokeapi.NewConfig()
	startRepl(&cfg)
}

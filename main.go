package main

import (
	"time"

	"github.com/Daiki1212/pokeDexki/Services/Pokeapi"
)

func main() {
	pokeClient := Pokeapi.NewClient(10 * time.Second)
	cfg := &config{pokeapiClient: pokeClient}
	startRepl(cfg)
}

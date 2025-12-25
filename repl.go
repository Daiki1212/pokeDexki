package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Daiki1212/pokeDexki/Application/Console/Commands"
	"github.com/Daiki1212/pokeDexki/Services/Pokeapi"
)

func startRepl(cfg *Pokeapi.Config, pokedex *map[string]Pokeapi.Pokemon) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		queryParams := strings.Join(words[1:], " ")

		command, exists := Commands.GetCommands()[commandName]
		if exists {
			err := command.Callback(cfg, queryParams, pokedex)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

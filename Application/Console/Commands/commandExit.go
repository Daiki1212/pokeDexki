package Commands

import (
	"fmt"
	"os"

	"github.com/Daiki1212/pokeDexki/Services/Pokeapi"
)

func CommandExit(_ *Pokeapi.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

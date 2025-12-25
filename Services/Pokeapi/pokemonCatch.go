package Pokeapi

import (
	"encoding/json"
)

func (c Client) GetPokemon(pokemonName string) (Pokemon, error) {
	url := baseUrl + pokemonPath + "/" + pokemonName

	body, err := c.getBodyFor(url)
	if err != nil {
		return Pokemon{}, err
	}

	pokemon := Pokemon{}
	if err = json.Unmarshal(body, &pokemon); err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}

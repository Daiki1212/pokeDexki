package Pokeapi

type Location struct {
	Name             string `json:"name"`
	Id               int    `json:"id"`
	PokemonEncounter []struct {
		Pokemon struct {
			Name string `json:"name"`
			Url  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

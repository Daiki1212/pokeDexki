package Pokeapi

import (
	"encoding/json"
)

func (c Client) GetLocation(location string) (Location, error) {
	url := baseUrl + locationPath + "/" + location

	body, err := c.getBodyFor(url)
	if err != nil {
		return Location{}, err
	}

	encounterList := Location{}
	if err := json.Unmarshal(body, &encounterList); err != nil {
		return Location{}, err
	}

	return encounterList, nil
}

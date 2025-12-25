package Pokeapi

import (
	"encoding/json"
)

func (c Client) ListLocations(pageUrl *string) (ResponseLocations, error) {
	url := baseUrl + locationPath
	if pageUrl != nil {
		url = *pageUrl
	}

	body, err := c.getBodyFor(url)
	if err != nil {
		return ResponseLocations{}, err
	}

	locationList := ResponseLocations{}
	if err := json.Unmarshal(body, &locationList); err != nil {
		return ResponseLocations{}, err
	}
	return locationList, nil
}

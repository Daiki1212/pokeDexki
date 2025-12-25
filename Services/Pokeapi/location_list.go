package Pokeapi

import (
	"encoding/json"
	"io"
)

func (c Client) ListLocations(pageUrl *string) (ResponseLocations, error) {
	url := baseUrl + locationPath
	if pageUrl != nil {
		url = *pageUrl
	}

	body, exists := c.cache.Get(url)
	if !exists {
		response, err := c.httpClient.Get(url)
		if err != nil {
			return ResponseLocations{}, err
		}
		defer response.Body.Close()

		body, err = io.ReadAll(response.Body)
		if err != nil {
			return ResponseLocations{}, err
		}

		c.cache.Add(url, body)
	}

	locationList := ResponseLocations{}
	if err := json.Unmarshal(body, &locationList); err != nil {
		return ResponseLocations{}, err
	}
	return locationList, nil
}

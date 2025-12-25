package Pokeapi

import (
	"encoding/json"
	"io"
)

func (c Client) GetLocation(location string) (Encounters, error) {
	url := baseUrl + locationPath + "/" + location

	body, exists := c.cache.Get(url)
	if !exists {
		response, err := c.httpClient.Get(url)
		if err != nil {
			return Encounters{}, err
		}
		defer response.Body.Close()

		body, err = io.ReadAll(response.Body)
		if err != nil {
			return Encounters{}, err
		}

		c.cache.Add(url, body)
	}

	encounterList := Encounters{}
	if err := json.Unmarshal(body, &encounterList); err != nil {
		return Encounters{}, err
	}

	return encounterList, nil
}

package Pokeapi

import "encoding/json"

func (c Client) List(pageUrl *string) (ResponseLocations, error) {
	url := baseUrl + locationPath
	if pageUrl != nil {
		url = *pageUrl
	}

	res, err := c.httpClient.Get(url)
	if err != nil {
		return ResponseLocations{}, err
	}
	defer res.Body.Close()

	locationList := ResponseLocations{}
	decoder := json.NewDecoder(res.Body)
	if err = decoder.Decode(&locationList); err != nil {
		return ResponseLocations{}, err
	}

	return locationList, nil
}

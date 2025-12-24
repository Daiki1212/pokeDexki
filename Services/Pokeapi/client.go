package Pokeapi

import (
	"net/http"
	"time"
)

type Config struct {
	PokeapiClient    Client
	NextLocationsURL *string
	PrevLocationsURL *string
}

type Client struct {
	httpClient http.Client
}

func NewConfig() Config {
	pokeClient := newClient(10 * time.Second)
	return Config{PokeapiClient: pokeClient}
}

func newClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}

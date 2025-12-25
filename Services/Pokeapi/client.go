package Pokeapi

import (
	"net/http"
	"time"

	"github.com/Daiki1212/pokeDexki/Services/Pokecache"
)

type Config struct {
	PokeapiClient    Client
	NextLocationsURL *string
	PrevLocationsURL *string
}

type Client struct {
	httpClient http.Client
	cache      *Pokecache.Cache
}

func NewConfig() Config {
	pokeClient := newClient(10 * time.Second)
	return Config{PokeapiClient: pokeClient}
}

func newClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{Timeout: timeout},
		cache:      Pokecache.NewCache(5 * time.Minute),
	}
}

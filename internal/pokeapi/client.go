package pokeapi

import (
	"net/http"
	"pokedex/internal/pokecache"
	"time"
)

type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

func NewClient(timeout time.Duration, interval time.Duration) Client {
	pokecache.NewCache(interval)
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}

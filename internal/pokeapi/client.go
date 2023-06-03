package pokeapi

import (
	"net/http"
	"time"

	"github.com/rjt5412/pokedexcli/internal/pokecache"
)

// Client
type Client struct {
	httpClient http.Client
	cache      pokecache.Cache
}

// New Client
func NewClient(cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}

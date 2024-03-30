package pokeapi

import (
	"net/http"
	"time"

	"github.com/codelikesuraj/pokedexcli/internal/pokecache"
)

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

func NewClient(cacheDuration time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheDuration),
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}

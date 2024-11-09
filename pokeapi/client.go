package pokeapi

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/mrkiz-git/pokedexcli/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2"

// Logger for the package
var logger = log.New(os.Stdout, "POKEAPI: ", log.Ltime)

// Client manages HTTP requests and caching for the PokeAPI.
type Client struct {
	httpClient *http.Client
	Cache      *pokecache.Cache
}

// NewClient creates and returns a new PokeAPI client.
func New() *Client {
	return &Client{
		httpClient: &http.Client{Timeout: 10 * time.Second},
		Cache:      pokecache.NewCache(5 * time.Minute),
	}
}

package pokeapi

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/mrkiz-git/pokedexcli/pokecache"
)

// Base URL for the PokeAPI
const baseURL = "https://pokeapi.co/api/v2"

// Package-level logger
var logger = log.New(os.Stdout, "POKEAPI: ", log.Ltime)

// Client represents a PokeAPI client with HTTP and caching capabilities
type Client struct {
	httpClient *http.Client
	cache      *pokecache.Cache
}

// New creates and returns a new PokeAPI client with caching enabled
func New() *Client {
	logger.Println("Creating new PokeAPI client with cache")
	return &Client{
		httpClient: &http.Client{},
		cache:      pokecache.NewCache(5 * time.Minute), // Cache entries expire after 5 minutes
	}
}

// ListLocationAreas returns a list of location areas from the PokeAPI
// If pageURL is provided, it will fetch that specific page
func (c *Client) ListLocationAreas(pageURL *string) (*ListResponse[LocationAreas], error) {
	// Determine the URL to use
	url := baseURL + "/location-area"
	if pageURL != nil {

		logger.Printf("Page Url: %s", *pageURL)
		url = *pageURL
	}
	logger.Printf("Processing request for URL: %s", url)

	// Check if the response is in cache
	if cachedData, ok := c.cache.Get(url); ok {
		logger.Printf("Cache hit for URL: %s", url)
		var response ListResponse[LocationAreas]
		err := json.Unmarshal(cachedData, &response)
		if err != nil {
			logger.Printf("Error unmarshaling cached data: %v", err)
			return nil, err
		}
		return &response, nil
	}

	// If not in cache, make the HTTP request
	logger.Printf("Cache miss - making HTTP request to: %s", url)
	resp, err := c.httpClient.Get(url)
	if err != nil {
		logger.Printf("Error making HTTP request: %v", err)
		return nil, err
	}
	defer resp.Body.Close()

	logger.Printf("Got response status: %s", resp.Status)

	// Decode the response
	var response ListResponse[LocationAreas]
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		logger.Printf("Error decoding response: %v", err)
		return nil, err
	}

	// Cache the successful response
	responseBytes, err := json.Marshal(response)
	if err != nil {
		logger.Printf("Error marshaling response for cache: %v", err)
		return nil, err
	}

	logger.Printf("Caching response for URL: %s", url)
	c.cache.Add(url, responseBytes)

	logger.Printf("Successfully processed response with %d results", len(response.Results))
	return &response, nil
}

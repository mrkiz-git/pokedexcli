package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Fetch fetches data from the PokeAPI and decodes it into the specified output type.
// It uses the provided Client for HTTP requests and caching.
func Fetch[T any](client *Client, url string, output *T) error {
	// Check cache first
	if cachedData, found := client.Cache.Get(url); found {
		if err := json.Unmarshal(cachedData, output); err == nil {
			logger.Printf("Cache hit for URL: %s", url)
			return nil
		} else {
			logger.Printf("Failed to unmarshal cached data for URL %s: %v", url, err)
		}
	}

	// Make HTTP request
	resp, err := client.httpClient.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch data from URL %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected HTTP status %d for URL %s", resp.StatusCode, url)
	}

	// Decode JSON response
	if err := json.NewDecoder(resp.Body).Decode(output); err != nil {
		return fmt.Errorf("failed to decode response from URL %s: %w", url, err)
	}

	// Cache the response
	responseBytes, err := json.Marshal(output)
	if err == nil {
		client.Cache.Add(url, responseBytes)
	}

	logger.Printf("Fetched and cached data for URL: %s", url)
	return nil
}

package pokeapi_test

import (
	"testing"

	"github.com/mrkiz-git/pokedexcli/pokeapi"
)

func TestCacheHit(t *testing.T) {
	client := pokeapi.New()

	// Preload cache
	url := "https://pokeapi.co/api/v2/location-area"
	cachedResponse := `{
		"count": 1,
		"next": null,
		"previous": null,
		"results": [
			{ "name": "area1", "url": "https://pokeapi.co/area1" }
		]
	}`
	client.Cache.Add(url, []byte(cachedResponse))

	// Call GetLocationAreas
	result, err := client.GetLocationAreas(nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Assert results
	if len(result.Results) != 1 {
		t.Fatalf("Expected 1 result, got %d", len(result.Results))
	}

	if result.Results[0].Name != "area1" {
		t.Errorf("Expected first result to be 'area1', got %s", result.Results[0].Name)
	}
}

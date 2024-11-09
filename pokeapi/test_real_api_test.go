package pokeapi_test

import (
	"testing"

	"github.com/mrkiz-git/pokedexcli/pokeapi"
)

func TestRealGetLocationAreas(t *testing.T) {
	client := pokeapi.New()

	result, err := client.GetLocationAreas(nil)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result.Count == 0 {
		t.Errorf("Expected non-zero count, got %d", result.Count)
	}
}

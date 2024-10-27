package pokeapi

import (
	"testing"
)

func TestListLocationAreas(t *testing.T) {
	client := New()

	resp, err := client.ListLocationAreas(nil)
	if err != nil {
		t.Fatalf("ListLocationAreas failed: %v", err)
	}

	// Print the full response using t.Log
	t.Logf("Full Response: %+v", resp)
	t.Logf("Count: %d", resp.Count)
	t.Logf("Results: %+v", resp.Results)

	if resp.Count == 0 {
		t.Error("Expected count to be greater than 0")
	}

	if len(resp.Results) == 0 {
		t.Error("Expected results to not be empty")
	}
}

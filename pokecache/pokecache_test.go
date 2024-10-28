package pokecache

import (
	"log"
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	log.Println("Testing cache creation...")
	cache := NewCache(time.Millisecond * 10)
	if cache == nil {
		t.Error("expected cache to not be nil")
	}
	log.Println("Cache creation test completed")
}

func TestAddGet(t *testing.T) {
	log.Println("Testing Add and Get operations...")

	cache := NewCache(time.Millisecond * 10)
	key := "test-key"
	val := []byte("test-value")

	log.Printf("Adding key: %s with value: %s", key, string(val))
	cache.Add(key, val)

	log.Printf("Attempting to retrieve key: %s", key)
	result, ok := cache.Get(key)
	if !ok {
		t.Error("expected to find key")
		return
	}

	if string(result) != string(val) {
		t.Errorf("expected to get %s, got %s", string(val), string(result))
	}
	log.Println("Add/Get test completed successfully")
}

func TestReap(t *testing.T) {
	log.Println("Testing reap functionality...")

	interval := time.Millisecond * 10
	log.Printf("Creating cache with interval: %v", interval)
	cache := NewCache(interval)

	key := "test-key"
	val := []byte("test-value")

	log.Printf("Adding key: %s with value: %s", key, string(val))
	cache.Add(key, val)

	log.Printf("Waiting %v for reap...", interval*2)
	time.Sleep(interval * 2)

	log.Printf("Checking if key %s was reaped", key)
	_, ok := cache.Get(key)
	if ok {
		t.Error("expected key to have been reaped")
	}
	log.Println("Reap test completed")
}

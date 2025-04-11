package main

import (
	"testing"

	table "github.com/muthu-kumar-u/go-hashmap/internal/table"
)

type testValue struct {
	Name string
}

// Helper to create a basic test value
func newTestValue(name string) *testValue {
	return &testValue{Name: name}
}

func TestAddAndGet(t *testing.T) {
	ht := table.NewHashTable(10)

	// Add one entry
	ht.Add(newTestValue("first"), "firstKey")

	// Retrieve it
	v, err := ht.Get("firstKey")
	if err != nil {
		t.Errorf("Expected to find key 'firstKey', but got error: %v", err)
	}
	if v.(*testValue).Name != "first" {
		t.Errorf("Expected value 'first', got %v", v)
	}

	// Try non-existent key
	_, err = ht.Get("unknown")
	if err != table.ErrKeyNotFound {
		t.Errorf("Expected ErrKeyNotFound for unknown key, got: %v", err)
	}
}

func TestUpdate(t *testing.T) {
	ht := table.NewHashTable(10)

	ht.Add(newTestValue("initial"), "key1")
	ok, err := ht.Update("key1", newTestValue("updated"))
	if err != nil || !ok {
		t.Errorf("Expected update to succeed, got error: %v", err)
	}

	v, _ := ht.Get("key1")
	if v.(*testValue).Name != "updated" {
		t.Errorf("Expected updated value 'updated', got: %v", v)
	}

	// Update non-existent key
	ok, err = ht.Update("missing", newTestValue("value"))
	if ok || err != table.ErrKeyNotFound {
		t.Errorf("Expected update to fail with ErrKeyNotFound, got: %v", err)
	}
}

func TestDelete(t *testing.T) {
	ht := table.NewHashTable(10)

	ht.Add(newTestValue("toDelete"), "keyToDelete")
	ok, err := ht.Delete("keyToDelete")
	if !ok || err != nil {
		t.Errorf("Expected successful deletion, got error: %v", err)
	}

	_, err = ht.Get("keyToDelete")
	if err != table.ErrKeyNotFound {
		t.Errorf("Expected key to be deleted, but still exists")
	}

	// Try deleting again
	ok, err = ht.Delete("keyToDelete")
	if ok || err != table.ErrKeyNotFound {
		t.Errorf("Expected delete to fail for non-existent key")
	}
}

func TestCollision(t *testing.T) {
	ht := table.NewHashTable(1) // Force collisions

	ht.Add(newTestValue("A"), "keyA")
	ht.Add(newTestValue("B"), "keyB")
	ht.Add(newTestValue("C"), "keyC")

	// All keys should still be retrievable
	keys := []string{"keyA", "keyB", "keyC"}
	for _, key := range keys {
		val, err := ht.Get(key)
		if err != nil || val == nil {
			t.Errorf("Expected to get key %s, got error: %v", key, err)
		}
	}
}

func TestGetAllEntry(t *testing.T) {
	ht := table.NewHashTable(3)

	_, err := ht.GetAllEntry()
	if err != nil {
		t.Errorf("Expected empty table to still return entries slice (with nil values)")
	}

	ht.Add(newTestValue("X"), "one")
	entries, err := ht.GetAllEntry()
	if err != nil {
		t.Errorf("Expected non-empty entries, got error: %v", err)
	}

	found := false
	for _, e := range entries {
		if e != nil && e.Value.(*testValue).Name == "X" {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Expected to find inserted entry in entries")
	}
}

package main

import (
	"fmt"

	table "github.com/muthu-kumar-u/go-hashmap/internal/table"
)

func main() {
	ht := table.NewHashTable(10)

	// Add entries
	value1 := map[string]string{"value 1": "add"}
	value2 := map[string]string{"value 2": "add"}

	ht.Add(value1, "firstKey")
	ht.Add(value1, "firstKey") // same key again (will chain, not update)
	ht.Add(value2, "secondKey")

	// Print all entries
	fmt.Println("\n--- After Adding ---")
	ht.Print()

	// Get entry
	_, err := ht.Get("firstKey")
	if err != nil {
		fmt.Println("Error getting key 'firstKey':", err.Error())
	}

	// Update entry
	updateValue := map[string]string{"update": "update"}
	ok, err := ht.Update("firstKey", updateValue)
	if err != nil || !ok {
		fmt.Println("Error updating key 'firstKey':", err.Error())
	}

	// Print all
	fmt.Println("\n--- After Update ---")
	ht.Print()

	// Handle collision
	collisionValue1 := map[string]string{"collision 1": "collision"}
	collisionValue2 := map[string]string{"collision 2": "collision"}

	ht.Add(collisionValue1, "thirdKey")
	ht.Add(collisionValue2, "fifthKey") // May collide, depending on hash

	// Print all entries
	fmt.Println("\n--- After Adding Collisions ---")
	ht.Print()

	// Delete entry
	_, err = ht.Delete("firstKey")
	if err != nil {
		fmt.Println("Error deleting key 'firstKey':", err.Error())
	}

	// Final print
	fmt.Println("\n--- After Deletion ---")
	ht.Print()
}

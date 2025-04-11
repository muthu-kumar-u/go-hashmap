package table

import (
	"errors"
	"fmt"
	"hash/fnv"

	"github.com/fatih/color"
)

var (
	ErrKeyNotFound = errors.New("key not found")
	ErrEmptyEntry  = errors.New("entries not found")
)

func NewHashTable(c int) *HashTable {
	return &HashTable{
		Entries: make([]*Entry, c),
		Size:    0,
	}
}

func (c *HashTable) Add(v interface{}, k string) {
	index := c.hash(k)
	entry := c.Entries[index]

	newEntry := &Entry{Key: k, Value: v, Next: nil}

	if entry == nil { // create new entry from the empty table
		c.Entries[index] = newEntry
		c.Size++ // increase the size
		return
	}

	// if not empty add the entry to the chain
	current := entry
	for current.Next != nil { // find the last next nil value (so that insert the value)
		current = current.Next // it will return the last nil valued entry
	}

	current.Next = newEntry // insert the entry to the nil entry
	c.Size++                // increase the size
	return
}

func (c *HashTable) Get(k string) (interface{}, error) {
	index := c.hash(k)
	current := c.Entries[index]

	for current != nil {
		if current.Key == k { // key founded
			return current.Value, nil
		}

		current = current.Next // move to the next node if the key not found
	}

	return nil, ErrKeyNotFound
}

func (c *HashTable) Delete(k string) (bool, error) {
	index := c.hash(k)
	current := c.Entries[index]
	var pre *Entry

	for current != nil {
		if current.Key == k {
			if pre == nil {
				c.Entries[index] = current.Next // head will removed directly
			} else {
				pre.Next = current.Next // remove the nth next entry from the chain
			}

			c.Size-- // decrease the size
			return true, nil
		}

		pre = current          // move the last visited node from chain
		current = current.Next // move to the next node from the chain
	}

	return false, ErrKeyNotFound
}

func (c *HashTable) Update(k string, v interface{}) (bool, error) {
	index := c.hash(k)
	current := c.Entries[index]

	for current != nil {
		if current.Key == k {
			current.Value = v // update the value if key matches
			return true, nil
		}
		current = current.Next // traverse to next node
	}

	return false, ErrKeyNotFound
}

func (c *HashTable) GetAllEntry() ([]*Entry, error) {
	if len(c.Entries) > 0 {
		return c.Entries, nil
	}

	return nil, ErrEmptyEntry
}

func (c *HashTable) Print() {
	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.BgYellow).SprintFunc()
	red := color.New(color.BgRed).SprintFunc()
	blue := color.New(color.BgBlue).SprintFunc()

	fmt.Printf("%s %d, %s %d\n", red("Total size:"), c.Size, blue("Capacity:"), len(c.Entries))

	for i, e := range c.Entries {
		if e != nil {
			fmt.Printf("%s %d: ", yellow("Bucket"), i+1)
			for n := e; n != nil; n = n.Next {
				fmt.Printf("%s -> ", green(fmt.Sprintf("[%s: %v]", n.Key, n.Value)))
			}
			fmt.Println("nil")
		}

	}
}

func (c *HashTable) hash(k string) uint {
	h := fnv.New32a()
	h.Write([]byte(k))
	return uint(h.Sum32()) % uint(len(c.Entries))
}

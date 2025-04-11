package table

type HashTable struct {
	Entries []*Entry
	Size    int
}

type Entry struct {
	Key   string
	Value interface{}
	Next  *Entry
}

package main

type KVS interface {
	// Inserts the given value into the store for the given key.
	Insert(key string, value int)

	// Returns the number of values in the store.
	Count() int

	// Returns a slice of all values for key if the given key is in
	// the store, and empty slice otherwise.
	// The order of returned values is not guaranteed.
	Search(key string) []int

	// Returns a slice of all values for key if key has the prefix
	// `prefix`, and empty slice otherwise.
	// The order of returned values is not guaranteed.
	PrefixSearch(prefix string) []int
}

func NewKVS() KVS {
	return NewMyKVS()
}

// --- Map KVS

type MyKVS struct {
}

func NewMyKVS() *MyKVS {
	return &MyKVS{}
}

func (m *MyKVS) Insert(key string, value int) {
}

func (m *MyKVS) Count() int {
	return 0
}

func (m *MyKVS) Search(key string) []int {
	return nil
}

func (m *MyKVS) PrefixSearch(prefix string) []int {
	return nil
}

package htable

type HashTable struct {
	data map[string]int
}

func NewHashTable() *HashTable {
	return &HashTable{data: make(map[string]int)}
}

func (ht *HashTable) Put(key string, value int) {
	ht.data[key] = value
}

func (ht *HashTable) Get(key string) (int, bool) {
	value, found := ht.data[key]
	return value, found
}

func (ht *HashTable) Delete(key string) {
	delete(ht.data, key)
}

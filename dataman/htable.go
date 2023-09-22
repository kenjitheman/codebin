package dataman

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

// func TryHT() {
// 	ht := NewHashTable()
// 	ht.Put("Alice", 25)
// 	ht.Put("Bob", 30)
//
// 	age, found := ht.Get("Alice")
// 	if found {
// 		fmt.Printf("Alice's age is %d\n", age)
// 	}
//
// 	ht.Delete("Bob")
// 	_, found = ht.Get("Bob")
// 	if !found {
// 		fmt.Println("Bob not found")
// 	}
// }

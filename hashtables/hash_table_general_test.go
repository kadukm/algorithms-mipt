package hashtables

import (
	"testing"
)

var (
	simpleHashFunc = func(s string) int {
		return len(s)
	}

	polynomialHashFunc = func(s string) (res int) {
		for i := 0; i < len(s); i++ {
			res = (res*239 + int(s[i])) % 1000000007
		}
		return
	}
)

type HashTable interface {
	Put(string, int)
	Get(string) (int, bool)
}

func TestPutAndGet(t *testing.T) {
	hashTables := createHashTables(polynomialHashFunc, 91)

	for _, hashTable := range hashTables {
		hashTable.Put("test", 10)
		value, ok := hashTable.Get("test")

		if !ok {
			t.Fatal(`value for key "test" should be present`)
		}
		if value != 10 {
			t.Fatalf(`value for key "test" should be equals to 10 but was %d`, value)
		}
	}
}

func TestGetLastEntryInBucket(t *testing.T) {
	hashTables := createHashTables(simpleHashFunc, 7)

	for _, hashTable := range hashTables {
		hashTable.Put("1st", 1)
		hashTable.Put("2nd", 2)
		hashTable.Put("3rd", 3)

		value, ok := hashTable.Get("3rd")

		if !ok {
			t.Fatal(`value for key "3rd" should be present`)
		}
		if value != 3 {
			t.Fatalf(`value for key "3rd" should be equals to 3 but was %d`, value)
		}
	}
}

func TestGetByNotExistingKeyFromNotEmptyBucket(t *testing.T) {
	hashTables := createHashTables(simpleHashFunc, 7)

	for _, hashTable := range hashTables {
		hashTable.Put("has", 100)

		value, ok := hashTable.Get("not")

		if ok {
			t.Fatalf(`value for key "not" should not be present but value was found: %d`, value)
		}
	}
}

func TestMultiplePutAndGet(t *testing.T) {
	hashTables := createHashTables(polynomialHashFunc, 91)

	entries := []struct {
		key   string
		value int
	}{
		{"a", 10},
		{"b", 100},
		{"c", 1000},
		{"abc", 123},
		{"heh", 322},
	}

	for _, hashTable := range hashTables {
		for _, e := range entries {
			hashTable.Put(e.key, e.value)
			value, ok := hashTable.Get(e.key)

			if !ok {
				t.Fatalf(`value for key "%s" should be present`, e.key)
			}
			if value != e.value {
				t.Fatalf(`value for key "%s" should be equals to %d but was %d`, e.key, e.value, value)
			}
		}
	}
}

func TestUpdateExistingEntry(t *testing.T) {
	hashTables := createHashTables(polynomialHashFunc, 91)

	for _, hashTable := range hashTables {
		hashTable.Put("test", 10)
		hashTable.Put("test", 100)
		value, ok := hashTable.Get("test")

		if !ok {
			t.Fatal(`value for key "test" should be present`)
		}
		if value != 100 {
			t.Fatalf(`value for key "test" should be equals to 100 but was %d`, value)
		}
	}
}

func TestUpdateSecondEntryFromBucket(t *testing.T) {
	hashTables := createHashTables(simpleHashFunc, 7)

	for _, hashTable := range hashTables {
		hashTable.Put("test1", 1)
		hashTable.Put("test2", 2)
		hashTable.Put("test2", 999)
		value, ok := hashTable.Get("test2")

		if !ok {
			t.Fatal(`value for key "test2" should be present`)
		}
		if value != 999 {
			t.Fatalf(`value for key "test2" should be equals to 999 but was %d`, value)
		}
	}
}

func TestGetByNotExistingKey(t *testing.T) {
	hashTables := createHashTables(polynomialHashFunc, 91)

	for _, hashTable := range hashTables {
		value, ok := hashTable.Get("not_existing_value")

		if ok {
			t.Fatalf(`value for key "not_existing_value" should not be present but value was found: %d`, value)
		}
	}
}

func TestPutKeysWithSameHashes(t *testing.T) {
	hashTables := createHashTables(simpleHashFunc, 7)

	for _, hashTable := range hashTables {
		hashTable.Put("a", 10)
		hashTable.Put("b", 100)

		value, ok := hashTable.Get("a")
		if !ok {
			t.Fatal(`value for key "a" should be present`)
		}
		if value != 10 {
			t.Fatalf(`value for key "a" should be equals to 10 but was %d`, value)
		}

		value, ok = hashTable.Get("b")
		if !ok {
			t.Fatal(`value for key "b" should be present`)
		}
		if value != 100 {
			t.Fatalf(`value for key "b" should be equals to 100 but was %d`, value)
		}
	}
}

func createHashTables(calcHashFunc func(string) int, size int) []HashTable {
	hashTables := make([]HashTable, 0, 2)

	hashTables = append(hashTables, NewSeparateChainingHashTable(calcHashFunc, size))
	hashTables = append(hashTables, createOpenAddressingHashTableAdapter(calcHashFunc, size))

	return hashTables
}

func createOpenAddressingHashTableAdapter(calcHashFunc func(string) int, size int) HashTable {
	hashTable := NewOpenAddressingHashTable(calcHashFunc, size)
	return &openAddressingHashTableAdapter{hashTable}
}

type openAddressingHashTableAdapter struct {
	hashTable OpenAddressingHashTable
}

func (adapter *openAddressingHashTableAdapter) Put(k string, v int) {
	if err := adapter.hashTable.Put(k, v); err != nil {
		panic(err)
	}
}

func (adapter *openAddressingHashTableAdapter) Get(k string) (int, bool) {
	return adapter.hashTable.Get(k)
}

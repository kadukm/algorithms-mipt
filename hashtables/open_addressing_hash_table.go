package hashtables

import "errors"

type OpenAddressingHashTable interface {
	Put(string, int) error
	Get(string) (int, bool)
}

type openAddressingHashTable struct {
	calcHash func(string) int
	buckets  []*simpleEntry
}

type simpleEntry struct {
	key   string
	value int
}

func NewOpenAddressingHashTable(calcHash func(string) int, size int) OpenAddressingHashTable {
	return &openAddressingHashTable{calcHash, make([]*simpleEntry, size)}
}

func (hashTable *openAddressingHashTable) Put(key string, value int) error {
	bucketIdx, entry := hashTable.getEntryOfBucket(key)

	if entry == nil {
		hashTable.buckets[bucketIdx] = &simpleEntry{key, value}
		return nil
	} else if entry.key == key {
		entry.value = value
		return nil
	}

	i := hashTable.calcNextIndex(bucketIdx)
	for {
		if i == bucketIdx {
			return errors.New("cannot put new key, hash table is full")
		}

		entry = hashTable.buckets[i]
		if entry == nil {
			hashTable.buckets[i] = &simpleEntry{key, value}
			return nil
		} else if entry.key == key {
			entry.value = value
			return nil
		}

		i = hashTable.calcNextIndex(i)
	}
}

func (hashTable *openAddressingHashTable) Get(key string) (int, bool) {
	bucketIdx, entry := hashTable.getEntryOfBucket(key)

	if entry == nil {
		return 0, false
	} else if entry.key == key {
		return entry.value, true
	}

	i := hashTable.calcNextIndex(bucketIdx)
	for {
		if i == bucketIdx {
			return 0, false
		}

		entry = hashTable.buckets[i]
		if entry == nil {
			return 0, false
		} else if entry.key == key {
			return entry.value, true
		}

		i = hashTable.calcNextIndex(i)
	}
}

func (hashTable *openAddressingHashTable) getEntryOfBucket(key string) (int, *simpleEntry) {
	h := hashTable.calcHash(key)
	i := h % len(hashTable.buckets)
	return i, hashTable.buckets[i]
}

func (hashTable *openAddressingHashTable) calcNextIndex(i int) int {
	return (i + 1) % len(hashTable.buckets)
}

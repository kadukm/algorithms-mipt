package hashtables

type SeparateChainingHashTable interface {
	Put(string, int)
	Get(string) (int, bool)
	Delete(string) bool
}

type separateChainingHashTable struct {
	calcHash func(string) int
	buckets  []*linkedEntry
}

type linkedEntry struct {
	key   string
	value int
	next  *linkedEntry
}

func NewSeparateChainingHashTable(calcHash func(string) int, size int) SeparateChainingHashTable {
	return &separateChainingHashTable{calcHash, make([]*linkedEntry, size)}
}

func (hashTable *separateChainingHashTable) Put(key string, value int) {
	i, entry := hashTable.getFirstEntryOfBucket(key)

	if entry == nil {
		hashTable.buckets[i] = &linkedEntry{key, value, nil}
		return
	}

	for {
		if entry.key == key {
			entry.value = value
			return
		}

		if entry.next == nil {
			break
		}

		entry = entry.next
	}

	entry.next = &linkedEntry{key, value, nil}
}

func (hashTable *separateChainingHashTable) Get(key string) (int, bool) {
	_, entry := hashTable.getFirstEntryOfBucket(key)

	if entry == nil {
		return 0, false
	}

	for entry != nil {
		if entry.key == key {
			return entry.value, true
		}

		entry = entry.next
	}

	return 0, false
}

func (hashTable *separateChainingHashTable) Delete(key string) bool {
	i, entry := hashTable.getFirstEntryOfBucket(key)

	if entry == nil {
		return false
	}

	if entry.key == key {
		if entry.next != nil {
			hashTable.buckets[i] = entry.next
		} else {
			hashTable.buckets[i] = nil
		}

		return true
	}

	prev := entry
	curr := prev.next
	for curr != nil {
		if curr.key == key {
			prev.next = curr.next
			return true
		}

		prev = curr
		curr = curr.next
	}

	return false
}

func (hashTable *separateChainingHashTable) getFirstEntryOfBucket(k string) (int, *linkedEntry) {
	h := hashTable.calcHash(k)
	i := h % len(hashTable.buckets)
	return i, hashTable.buckets[i]
}

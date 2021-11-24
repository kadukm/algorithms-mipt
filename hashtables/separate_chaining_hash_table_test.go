package hashtables

import "testing"

func TestDeleteExistingEntry(t *testing.T) {
	hashTable := NewSeparateChainingHashTable(polynomialHashFunc, 91)
	hashTable.Put("test", 10)

	ok := hashTable.Delete("test")

	if !ok {
		t.Fatal(`key "test" should be deleted successfully`)
	}

	value, ok := hashTable.Get("test")
	if ok {
		t.Fatalf(`key "test" should be deleted but value %d was found`, value)
	}
}

func TestDeleteByNotExistingKeyFromBucketWithMultipleValues(t *testing.T) {
	hashTable := NewSeparateChainingHashTable(simpleHashFunc, 7)
	hashTable.Put("has1", 1)
	hashTable.Put("has2", 2)

	ok := hashTable.Delete("test")

	if ok {
		t.Fatal(`key "test" should not be present but it was deleted successfully`)
	}
}

func TestDeleteFirstEntryFromBucketWithMultipleEntries(t *testing.T) {
	hashTable := NewSeparateChainingHashTable(simpleHashFunc, 7)
	hashTable.Put("1st", 1)
	hashTable.Put("2nd", 2)

	ok := hashTable.Delete("1st")

	if !ok {
		t.Fatal(`key "1st" should be deleted successfully`)
	}

	value, ok := hashTable.Get("2nd")
	if !ok {
		t.Fatal(`value for key "2nd" should be present`)
	}
	if value != 2 {
		t.Fatalf(`value for key "2nd" should be equals to 2 but was %d`, value)
	}
}

func TestDeleteLastEntryFromBucket(t *testing.T) {
	hashTable := NewSeparateChainingHashTable(simpleHashFunc, 7)
	hashTable.Put("1st", 1)
	hashTable.Put("2nd", 2)
	hashTable.Put("3rd", 3)
	hashTable.Put("4th", 4)

	ok := hashTable.Delete("4th")

	if !ok {
		t.Fatal(`key "4th" should be deleted successfully`)
	}

	value, ok := hashTable.Get("2nd")
	if !ok {
		t.Fatal(`value for key "2nd" should be present`)
	}
	if value != 2 {
		t.Fatalf(`value for key "2nd" should be equals to 2 but was %d`, value)
	}
}

func TestDeleteByNotExistingKey(t *testing.T) {
	hashTable := NewSeparateChainingHashTable(polynomialHashFunc, 91)

	ok := hashTable.Delete("test")

	if ok {
		t.Fatal(`key "test" should not be present but it was deleted successfully`)
	}
}

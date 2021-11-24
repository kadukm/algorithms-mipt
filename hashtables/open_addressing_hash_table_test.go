package hashtables

import "testing"

func TestFillAddBuckets(t *testing.T) {
	hashTable := NewOpenAddressingHashTable(simpleHashFunc, 3)

	for _, k := range []string{"1", "2", "33"} {
		err := hashTable.Put(k, 10)

		if err != nil {
			t.Fatalf(`key "%s" should be but successfully but was error: %s`, k, err)
		}
	}
}

func TestPutMoreEntriesThanHashTableSize(t *testing.T) {
	hashTable := NewOpenAddressingHashTable(simpleHashFunc, 3)

	for _, k := range []string{"1", "2", "33"} {
		err := hashTable.Put(k, 10)
		if err != nil {
			t.Fatalf(`key "%s" should be but successfully but was error: %s`, k, err)
		}
	}

	err := hashTable.Put("error_key", 1234)
	if err == nil {
		t.Fatal(`last key "error_key" should not be put in hash table`)
	}
}

func TestGetByNotExistingKeyWhenAllBucketsFilled(t *testing.T) {
	hashTable := NewOpenAddressingHashTable(simpleHashFunc, 3)

	for _, k := range []string{"1", "2", "33"} {
		err := hashTable.Put(k, 10)
		if err != nil {
			t.Fatalf(`key "%s" should be put successfully but was error: %s`, k, err)
		}
	}

	if value, ok := hashTable.Get("not_existing_key"); ok {
		t.Fatalf(`value for key "not" should not be present but value was found: %d`, value)
	}
}

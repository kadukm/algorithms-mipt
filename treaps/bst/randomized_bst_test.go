package bst

import "testing"

func TestAddValuesInRandomOrder(t *testing.T) {
	expectedResult := []int{-1, 0, 6, 9, 17}
	treap := new(Treap)

	for _, index := range []int{0, 2, 1, 4, 3} {
		treap.Add(expectedResult[index])
	}

	actualResult := treap.GetSortedValues()
	if !slicesEqual(expectedResult, actualResult) {
		t.Fatalf("expected %v but was %v", expectedResult, actualResult)
	}
}

func TestAddOrderedValues(t *testing.T) {
	expectedResult := []int{-1, 0, 6, 9, 17}
	treap := new(Treap)

	for _, value := range expectedResult {
		treap.Add(value)
	}

	actualResult := treap.GetSortedValues()
	if !slicesEqual(expectedResult, actualResult) {
		t.Fatalf("expected %v but was %v", expectedResult, actualResult)
	}
}

func TestAddSameValues(t *testing.T) {
	treap := new(Treap)
	expectedResult := make([]int, 100)
	for i := 0; i < 100; i++ {
		expectedResult[i] = 99
	}

	for _, value := range expectedResult {
		treap.Add(value)
	}

	actualResult := treap.GetSortedValues()
	if !slicesEqual(expectedResult, actualResult) {
		t.Fatalf("expected %v but was %v", expectedResult, actualResult)
	}
}

func TestContainsInEmptyTreap(t *testing.T) {
	treap := new(Treap)

	ok := treap.Contains(1337)

	if ok {
		t.Fatalf("value 1337 doesn't exist in the treap but it was found")
	}
}

func TestContainsInTreapWithOneValue(t *testing.T) {
	treap := new(Treap)
	treap.Add(9)

	ok := treap.Contains(9)

	if !ok {
		t.Fatalf("value 9 exists in the treap but it wasn't found")
	}
}

func TestContainsInTreapWithMultipleValues(t *testing.T) {
	treap := new(Treap)
	values := []int{-9, 600, 300, 1, 0, 10, 45, 17, 365, 366, 11, 1998}

	for _, value := range values {
		treap.Add(value)
	}

	for _, value := range values {
		ok := treap.Contains(value)

		if !ok {
			t.Fatalf("value %d exists in the treap but it wasn't found", value)
		}
	}
}

func TestContainsInTreapWithSameValues(t *testing.T) {
	treap := new(Treap)
	for i := 0; i < 100; i++ {
		treap.Add(-765)
	}

	ok := treap.Contains(-765)

	if !ok {
		t.Fatalf("value -765 exists in the treap but it wasn't found")
	}
}

func TestDeleteFromEmptyTreap(t *testing.T) {
	treap := new(Treap)

	ok := treap.Delete(1337)

	if ok {
		t.Fatalf("value 1337 doesn't exist in the treap but it was deleted successfully")
	}
}

func TestDeleteFromTreapWithOneValue(t *testing.T) {
	treap := new(Treap)
	treap.Add(999)

	ok := treap.Delete(999)

	if !ok {
		t.Fatalf("value 999 should be deleted successfully but it wasn't")
	}
	if treap.Contains(999) {
		t.Fatalf("value 999 should be deleted but it is still in the treap")
	}
}

func TestDeleteAllValuesFromTreap(t *testing.T) {
	treap := new(Treap)
	values := []int{123, 34523, -12, 34, 0, -2344, 1, 235599, -1111, 2021, -2022}

	for _, value := range values {
		treap.Add(value)
	}

	for _, value := range values {
		ok := treap.Delete(value)

		if !ok {
			t.Fatalf("value %d should be deleted successfully but it wasn't", value)
		}
		if treap.Contains(value) {
			t.Fatalf("value %d should be deleted but it is still in the treap", value)
		}
	}
}

func TestDeleteValueWhichPresentedInTheTripMultipleTimes(t *testing.T) {
	treap := new(Treap)
	for i := 0; i < 100; i++ {
		treap.Add(159)
	}

	ok := treap.Delete(159)

	if !ok {
		t.Fatalf("value 159 should be deleted successfully but it wasn't")
	}
	if treap.Contains(159) {
		t.Fatalf("value 159 should be deleted but it is still in the treap")
	}
}

func slicesEqual(a1, a2 []int) bool {
	if len(a1) != len(a2) {
		return false
	}

	for i := 0; i < len(a1); i++ {
		if a1[i] != a2[i] {
			return false
		}
	}

	return true
}

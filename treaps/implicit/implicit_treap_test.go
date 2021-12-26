package implicit

import "testing"

func TestGetFromEmptyTreap(t *testing.T) {
	treap := new(ImplicitTreap)

	value, err := treap.Get(0)

	if err == nil {
		t.Fatalf(`expected "out of range" error but got value %d`, value)
	}
}

func TestAddAndGet(t *testing.T) {
	treap := new(ImplicitTreap)

	treap.Add(9)
	value, err := treap.Get(0)

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if value != 9 {
		t.Fatalf("expected value 9 but was %d", value)
	}
}

func TestAddAndGetFewValues(t *testing.T) {
	treap := new(ImplicitTreap)
	values := []int{9, -1, 11, 11, 16, -99, 10, 0}

	for _, value := range values {
		treap.Add(value)
	}
	for i, expectedValue := range values {
		actualValue, err := treap.Get(i)

		if err != nil {
			t.Fatalf("unexpected error on getting value on index %d: %s", i, err)
		}
		if actualValue != expectedValue {
			t.Fatalf("expected value %d but was %d", expectedValue, actualValue)
		}
	}
}

func TestInsertAndGetFewValues(t *testing.T) {
	treap := new(ImplicitTreap)
	pairs := []struct {
		index int
		value int
	}{
		{0, 60},
		{0, 40},
		{2, -11},
		{1, 0},
		{3, 4},
	}

	for _, pair := range pairs {
		err := treap.Insert(pair.value, pair.index)
		if err != nil {
			t.Fatalf("unexpected error on inserting value %d: %s", pair.value, err)
		}
	}

	for i, expectedValue := range []int{40, 0, 60, 4, -11} {
		actualValue, err := treap.Get(i)

		if err != nil {
			t.Fatalf("unexpected error on getting value on index %d: %s", i, err)
		}
		if actualValue != expectedValue {
			t.Fatalf("expected value %d but was %d", expectedValue, actualValue)
		}
	}
}

func TestInsertOutOfRangeInEmptyTreap(t *testing.T) {
	treap := new(ImplicitTreap)

	err := treap.Insert(600, 1)

	if err == nil {
		t.Fatalf(`expected "out of range" error but value was inserted`)
	}
}

func TestInsertWithNegativeIndex(t *testing.T) {
	treap := new(ImplicitTreap)
	treap.Add(10)

	err := treap.Insert(-10, -1)

	if err == nil {
		t.Fatalf(`expected "out of range" error but value was inserted`)
	}
}

func TestInsertOutOfRangeInNotEmptyTreap(t *testing.T) {
	treap := new(ImplicitTreap)
	treap.Add(123)
	treap.Add(13)

	err := treap.Insert(3, 3)

	if err == nil {
		t.Fatalf(`expected "out of range" error but value was inserted`)
	}
}

func TestDeleteOnlyOneValue(t *testing.T) {
	treap := new(ImplicitTreap)
	treap.Add(6)

	err := treap.Delete(0)

	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	_, err = treap.Get(0)
	if err == nil {
		t.Fatal("expected to delete value 6 but it's still in the treap")
	}
}

func TestDeleteAllValues(t *testing.T) {
	treap := new(ImplicitTreap)
	for _, value := range []int{10, 16, 25, 34, 665, -1, -22, 150} {
		treap.Add(value)
	}

	for i := 0; i < 8; i++ {
		err := treap.Delete(0)

		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
	}

	_, err := treap.Get(0)
	if err == nil {
		t.Fatal("expected to delete value 6 but it's still in the treap")
	}
}

func TestDeleteFromEmptyTreap(t *testing.T) {
	treap := new(ImplicitTreap)

	err := treap.Delete(0)

	if err == nil {
		t.Fatal("treap was empty but some element was deleted")
	}
}

func TestDeleteOnNegativeIndex(t *testing.T) {
	treap := new(ImplicitTreap)

	err := treap.Delete(-1)

	if err == nil {
		t.Fatal(`expected "out of range" error but some element was deleted`)
	}
}

func TestDeleteOutOfRange(t *testing.T) {
	treap := new(ImplicitTreap)
	treap.Add(123)
	treap.Add(987)

	err := treap.Delete(3)

	if err == nil {
		t.Fatal(`expected "out of range" error but some element was deleted`)
	}
}

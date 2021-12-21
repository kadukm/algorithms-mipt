package heaps

import "testing"

type namedHeap struct {
	name string
	heap Heap
}

func TestGetMinFromEmptyHeap(t *testing.T) {
	namedHeaps := createNamedHeaps()

	for _, currentNamedHeap := range namedHeaps {
		value, err := currentNamedHeap.heap.GetMin()

		if err == nil {
			t.Fatalf("%s: expected error but was value %d", currentNamedHeap.name, value)
		}
	}
}

func TestPopFromEmptyHeap(t *testing.T) {
	namedHeaps := createNamedHeaps()

	for _, currentNamedHeap := range namedHeaps {
		value, err := currentNamedHeap.heap.Pop()

		if err == nil {
			t.Fatalf("%s: expected error but was value %d", currentNamedHeap.name, value)
		}
	}
}

func TestAddAndPop(t *testing.T) {
	namedHeaps := createNamedHeaps()

	for _, currentNamedHeap := range namedHeaps {
		currentNamedHeap.heap.Add(6)

		value, err := currentNamedHeap.heap.Pop()

		if err != nil {
			t.Fatalf("%s: unexpected error: %s", currentNamedHeap.name, err)
		}
		if value != 6 {
			t.Fatalf("%s: expected 6 but was %d", currentNamedHeap.name, value)
		}
	}
}

func TestAddAndGetMin(t *testing.T) {
	namedHeaps := createNamedHeaps()

	for _, currentNamedHeap := range namedHeaps {
		currentNamedHeap.heap.Add(-5)

		value, err := currentNamedHeap.heap.GetMin()

		if err != nil {
			t.Fatalf("%s: unexpected error: %s", currentNamedHeap.name, err)
		}
		if value != -5 {
			t.Fatalf("%s: expected -5 but was %d", currentNamedHeap.name, value)
		}
	}
}

func TestFewAddAndGetMin(t *testing.T) {
	namedHeaps := createNamedHeaps()

	for _, currentNamedHeap := range namedHeaps {
		for _, value := range []int{1, -5, 10, 16, -100, 0} {
			currentNamedHeap.heap.Add(value)
		}

		value, err := currentNamedHeap.heap.GetMin()

		if err != nil {
			t.Fatalf("%s: unexpected error: %s", currentNamedHeap.name, err)
		}
		if value != -100 {
			t.Fatalf("%s: expected -100 but was %d", currentNamedHeap.name, value)
		}
	}
}

func TestFewAddAndPop(t *testing.T) {
	namedHeaps := createNamedHeaps()

	for _, currentNamedHeap := range namedHeaps {
		for _, value := range []int{600, 500, 4, 5, 0} {
			currentNamedHeap.heap.Add(value)
		}

		value, err := currentNamedHeap.heap.Pop()

		if err != nil {
			t.Fatalf("%s: unexpected error: %s", currentNamedHeap.name, err)
		}
		if value != 0 {
			t.Fatalf("%s: expected -0 but was %d", currentNamedHeap.name, value)
		}
	}
}

func TestFewAddAndCleanHeap(t *testing.T) {
	namedHeaps := createNamedHeaps()

	for _, currentNamedHeap := range namedHeaps {
		for _, value := range []int{10, 10, 10, 0, -1, 15, -9, 100} {
			currentNamedHeap.heap.Add(value)
		}

		for _, expectedValue := range []int{-9, -1, 0, 10, 10, 10, 15, 100} {
			actualValue, err := currentNamedHeap.heap.Pop()

			if err != nil {
				t.Fatalf("%s: unexpected error: %s", currentNamedHeap.name, err)
			}
			if actualValue != expectedValue {
				t.Fatalf("%s: expected %d but was %d", currentNamedHeap.name, expectedValue, actualValue)
			}
		}

	}
}

func TestFewAddThenCleanHeapAndGetMin(t *testing.T) {
	namedHeaps := createNamedHeaps()

	for _, currentNamedHeap := range namedHeaps {
		for _, value := range []int{6, 10, -1, 0, 300, 289, -100, 4, 7, -9} {
			currentNamedHeap.heap.Add(value)
		}

		for i := 0; i < 10; i++ {
			_, err := currentNamedHeap.heap.Pop()

			if err != nil {
				t.Fatalf("%s: unexpected error: %s", currentNamedHeap.name, err)
			}
		}

		value, err := currentNamedHeap.heap.GetMin()

		if err == nil {
			t.Fatalf("%s: expected error but was value %d", currentNamedHeap.name, value)
		}
	}
}

func TestFewAddThenCleanHeapAndPop(t *testing.T) {
	namedHeaps := createNamedHeaps()

	for _, currentNamedHeap := range namedHeaps {
		for _, value := range []int{11, -11, 0, 600, 1111, 55, 64, -77, 16, 11} {
			currentNamedHeap.heap.Add(value)
		}

		for i := 0; i < 10; i++ {
			_, err := currentNamedHeap.heap.Pop()

			if err != nil {
				t.Fatalf("%s: unexpected error: %s", currentNamedHeap.name, err)
			}
		}

		value, err := currentNamedHeap.heap.Pop()

		if err == nil {
			t.Fatalf("%s: expected error but was value %d", currentNamedHeap.name, value)
		}
	}
}

func createNamedHeaps() []namedHeap {
	return []namedHeap{
		{"BinaryHeap", NewBinaryHeap()},
		{"RandomizedHeap", NewRandomizedHeap()},
	}
}

package stack

import "testing"

func TestCapacityIncreasedOnPush(t *testing.T) {
	s := stack{make([]int, 0, 1)}

	s.Push(10)
	s.Push(20)

	if cap(s.data) != 2 {
		t.Fatalf("capacity was expected to be 2 but it was %d", cap(s.data))
	}
}

func TestCapacityNotChangedOnPop(t *testing.T) {
	s := stack{make([]int, 0, 10)}
	for i := 0; i < 10; i++ {
		s.Push(i)
	}

	for i := 0; i < 10; i++ {
		_, _ = s.Pop()
	}

	if cap(s.data) != 10 {
		t.Fatalf("capacity was expected to be 10 but it was %d", cap(s.data))
	}
}

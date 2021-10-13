package stack

import "testing"

const (
	expectedPopErrorMessage  = "attempt to pop from empty stack is incorrect"
	expectedPeekErrorMessage = "attempt to peek from empty stack is incorrect"
)

func TestPushMoreThanInitialCapacity(t *testing.T) {
	s := NewStack(1)

	s.Push(10)
	s.Push(20)
}

func TestPopOnFilledStack(t *testing.T) {
	s := NewStack(10)
	s.Push(123)

	value, err := s.Pop()

	if err != nil {
		t.Fatalf("error wasn't expected on pop but it was: %s", err)
	}
	if value != 123 {
		t.Fatalf("popped value was expected to be 123 but it was %d", value)
	}
}

func TestPopOnEmptyStack(t *testing.T) {
	s := NewStack(10)

	_, err := s.Pop()

	if err == nil {
		t.Fatal("error was expected on pop but it wasn't")
	}
	if err.Error() != expectedPopErrorMessage {
		t.Fatalf(`error was expected to be "%s" on pop, but it was "%s"`, expectedPopErrorMessage, err.Error())
	}
}

func TestPeekOnFilledStack(t *testing.T) {
	s := NewStack(10)
	s.Push(321)

	value, err := s.Peek()

	if err != nil {
		t.Fatalf("error wasn't expected on peek but it was: %s", err)
	}
	if value != 321 {
		t.Fatalf("peeked value was expected to be 123 but it was %d", value)
	}
}

func TestPeekOnEmptyStack(t *testing.T) {
	s := NewStack(10)

	_, err := s.Peek()

	if err == nil {
		t.Fatal("error was expected on peek but it wasn't")
	}
	if err.Error() != expectedPeekErrorMessage {
		t.Fatalf(`error was expected to be "%s" on peek, but it was "%s"`, expectedPeekErrorMessage, err.Error())
	}
}

func TestSizeOnFilledStack(t *testing.T) {
	s := NewStack(10)
	for i := 0; i < 5; i++ {
		s.Push(i)
	}

	size := s.Size()

	if size != 5 {
		t.Fatalf("size was expected to be 5 but it was %d", size)
	}
}

func TestSizeOnEmptyStack(t *testing.T) {
	s := NewStack(10)

	size := s.Size()

	if size != 0 {
		t.Fatalf("size was expected to be empty but it was %d", size)
	}
}

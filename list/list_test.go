package list

import "testing"

func TestAddLeftElementInEmptyListThenGetRightElement(t *testing.T) {
	l := NewList()

	l.AddLeft(-10)
	value, err := l.GetRight()

	if err != nil {
		t.Fatalf("error wasn't expected on getting right but it was: %s", err)
	}
	if value != -10 {
		t.Fatalf("value was expected to be -10 but it was %d", value)
	}
}

func TestAddRightElementInEmptyListThenGetLeftElement(t *testing.T) {
	l := NewList()

	l.AddRight(10)
	value, err := l.GetLeft()

	if err != nil {
		t.Fatalf("error wasn't expected on getting left but it was: %s", err)
	}
	if value != 10 {
		t.Fatalf("value was expected to be 10 but it was %d", value)
	}
}

func TestRemoveAllElementsStartingFromLeft(t *testing.T) {
	l := NewList()
	for i := 0; i < 10; i++ {
		l.AddLeft(i)
	}

	for i := 0; i < 10; i++ {
		_, _ = l.RemoveLeft()
	}
	_, errLeft := l.GetLeft()
	_, errRight := l.GetRight()

	if errLeft == nil {
		t.Error("error was expected on getting left element but it wasn't")
	}
	if errRight == nil {
		t.Error("error was expected on getting right element but it wasn't")
	}
}

func TestRemoveAllElementsFromRight(t *testing.T) {
	l := NewList()
	for i := 0; i < 10; i++ {
		l.AddRight(i)
	}

	for i := 0; i < 10; i++ {
		_, _ = l.RemoveRight()
	}
	_, errLeft := l.GetLeft()
	_, errRight := l.GetRight()

	if errLeft == nil {
		t.Error("error was expected on getting left element but it wasn't")
	}
	if errRight == nil {
		t.Error("error was expected on getting right element but it wasn't")
	}
}

func TestRemoveRightFromEmptyList(t *testing.T) {
	l := NewList()

	_, err := l.RemoveRight()

	if err == nil {
		t.Fatal("error was expected on removing right element but it wasn't")
	}
}

func TestRemoveLeftFromEmptyList(t *testing.T) {
	l := NewList()

	_, err := l.RemoveLeft()

	if err == nil {
		t.Fatal("error was expected on removing left element but it wasn't")
	}
}

func TestAddElementsBySidesThenEmptyListBySides(t *testing.T) {
	l := NewList()

	l.AddLeft(20)
	l.AddRight(80)
	valueLeft, errLeft := l.RemoveLeft()
	valueRight, errRight := l.RemoveRight()

	if errLeft != nil {
		t.Fatalf("error wasn't expected on removing left element it was: %s", errLeft)
	}
	if valueLeft != 20 {
		t.Fatalf("left value was expected to be 20 but it was %d", valueLeft)
	}

	if errRight != nil {
		t.Fatalf("error wasn't expected on removing right element but it was: %s", errRight)
	}
	if valueRight != 80 {
		t.Fatalf("right value was expected to be 80 but it was %d", valueRight)
	}
}

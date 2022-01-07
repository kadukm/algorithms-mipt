package disjointset

import "testing"

func TestMakeSetForSameElementTwice(t *testing.T) {
	set := NewDisjointSet()

	err := set.MakeSet(9)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	err = set.MakeSet(9)
	if err == nil {
		t.Fatal("error expected but it wasn't")
	}
}

func TestFindElementWhichNotInSet(t *testing.T) {
	set := NewDisjointSet()

	representative, err := set.Find(1444)
	if err == nil {
		t.Fatalf("error expected but representative value %d was found", representative)
	}
}

func TestUnionElementsNotInSet(t *testing.T) {
	set := NewDisjointSet()

	err := set.Union(123, 34)
	if err == nil {
		t.Fatal("error expected but it wasn't")
	}
}

func TestUnionElementWithElementWhichNotInSet(t *testing.T) {
	set := NewDisjointSet()
	_ = set.MakeSet(0)

	err := set.Union(0, 89)
	if err == nil {
		t.Fatal("error expected but it wasn't")
	}

	err = set.Union(89, 0)
	if err == nil {
		t.Fatal("error expected but it wasn't")
	}
}

func TestMakeSetAndThenFind(t *testing.T) {
	set := NewDisjointSet()
	_ = set.MakeSet(78)

	representative, err := set.Find(78)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if representative != 78 {
		t.Fatalf("representative for element 78 was expected to be 78 but was %d", representative)
	}
}

func TestAfterUnionTwoElementsHaveSameRepresentative(t *testing.T) {
	set := NewDisjointSet()
	_ = set.MakeSet(1)
	_ = set.MakeSet(2)

	err := set.Union(1, 2)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	representative1, err := set.Find(1)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	representative2, err := set.Find(2)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if representative1 != representative2 {
		t.Fatalf("representatives for elements 1 and 2 expected to be equeals but was %d and %d",
			representative1, representative2)
	}
}

func TestOnUnionRepresentativeIsSelectedFromSetWithMoreElementsCount1(t *testing.T) {
	set := NewDisjointSet()
	_ = set.MakeSet(10)
	_ = set.MakeSet(20)
	_ = set.Union(10, 20)
	_ = set.MakeSet(30)
	expectedRepresentative, _ := set.Find(10)

	err := set.Union(10, 30)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	actualRepresentative, _ := set.Find(30)
	if actualRepresentative != expectedRepresentative {
		t.Fatalf("expected representative is %d but was %d", expectedRepresentative, actualRepresentative)
	}
}

func TestOnUnionRepresentativeIsSelectedFromSetWithMoreElementsCount2(t *testing.T) {
	set := NewDisjointSet()
	_ = set.MakeSet(10)
	_ = set.MakeSet(20)
	_ = set.Union(10, 20)
	_ = set.MakeSet(30)
	expectedRepresentative, _ := set.Find(10)

	err := set.Union(30, 20)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	actualRepresentative, _ := set.Find(30)
	if actualRepresentative != expectedRepresentative {
		t.Fatalf("expected representative is %d but was %d", expectedRepresentative, actualRepresentative)
	}
}

func TestAllElementsFromSameSetHaveSameRepresentative(t *testing.T) {
	set := NewDisjointSet()
	for i := 0; i < 100; i++ {
		_ = set.MakeSet(i)
	}
	for i := 1; i < 100; i++ {
		_ = set.Union(i-1, i)
	}

	expectedRepresentative, _ := set.Find(0)
	for i := 1; i < 100; i++ {
		representative, _ := set.Find(i)
		if representative != expectedRepresentative {
			t.Fatalf("representatives for all elements expected to be equeals but was %d and %d",
				representative, expectedRepresentative)
		}
	}
}

func TestUnionSameElements(t *testing.T) {
	set := NewDisjointSet()
	_ = set.MakeSet(666)

	err := set.Union(666, 666)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	representative, err := set.Find(666)
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if representative != 666 {
		t.Fatalf("representative for element 666 was expected to be 666 but was %d", representative)
	}
}

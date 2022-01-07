package segmenttree

import "testing"

func TestGetSumOfAllElements(t *testing.T) {
	tree := NewSegmentTree([]int{1, -2, 3, 4, -5, 6, 7, 8, -9})

	sum := tree.GetSum(0, 9)

	if sum != 13 {
		t.Fatalf("expected sum is 45 but was %d", sum)
	}
}

func TestGetAllPossibleSums(t *testing.T) {
	tree := NewSegmentTree([]int{1, 2, 3, 4})

	testCases := []struct {
		findL, findR, expectedSum int
	}{
		{0, 4, 10},
		{0, 3, 6},
		{1, 4, 9},
		{0, 2, 3},
		{1, 3, 5},
		{2, 4, 7},
		{0, 1, 1},
		{1, 2, 2},
		{2, 3, 3},
		{3, 4, 4},
	}

	for _, testCase := range testCases {
		actualSum := tree.GetSum(testCase.findL, testCase.findR)

		if actualSum != testCase.expectedSum {
			t.Fatalf("expected sum for segment [%d, %d) is %d but was %d",
				testCase.findL, testCase.findR, testCase.expectedSum, actualSum)
		}
	}
}

func TestGetAllPossibleSumsAfterSomeElementsChanged(t *testing.T) {
	tree := NewSegmentTree([]int{1, 2, 3, 4})
	if err := tree.Set(0, 10); err != nil {
		t.Fatalf("unexpected error on setting value 10 on index 0: %s", err)
	}
	if err := tree.Set(1, 25); err != nil {
		t.Fatalf("unexpected error on setting value 25 on index 1: %s", err)
	}

	testCases := []struct {
		findL, findR, expectedSum int
	}{
		{0, 4, 42},
		{0, 3, 38},
		{1, 4, 32},
		{0, 2, 35},
		{1, 3, 28},
		{2, 4, 7},
		{0, 1, 10},
		{1, 2, 25},
		{2, 3, 3},
		{3, 4, 4},
	}

	for _, testCase := range testCases {
		actualSum := tree.GetSum(testCase.findL, testCase.findR)

		if actualSum != testCase.expectedSum {
			t.Fatalf("expected sum for segment [%d, %d) is %d but was %d",
				testCase.findL, testCase.findR, testCase.expectedSum, actualSum)
		}
	}
}

func TestGetSumOfOutOfRangeSegments(t *testing.T) {
	tree := NewSegmentTree([]int{1, 2, 3, 4})

	testCases := []struct {
		findL, findR, expectedSum int
	}{
		{-10, 100, 10},
		{-135, 3, 6},
		{1, 12346, 9},
	}

	for _, testCase := range testCases {
		actualSum := tree.GetSum(testCase.findL, testCase.findR)

		if actualSum != testCase.expectedSum {
			t.Fatalf("expected sum for segment [%d, %d) is %d but was %d",
				testCase.findL, testCase.findR, testCase.expectedSum, actualSum)
		}
	}
}

func TestSetWithNegativeIndex(t *testing.T) {
	tree := NewSegmentTree([]int{1})

	err := tree.Set(-1, 123)

	if err == nil {
		t.Fatalf("expected error but it wasn't")
	}
}

func TestSetOutOfRange(t *testing.T) {
	tree := NewSegmentTree([]int{0, -1, 16})

	err := tree.Set(3, 123)

	if err == nil {
		t.Fatalf("expected error but it wasn't")
	}
}

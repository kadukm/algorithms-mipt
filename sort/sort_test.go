package sort

import (
	"testing"
)

func TestSortFunctions(t *testing.T) {
	sortFunctions := []struct {
		sort func([]int) []int
		name string
	}{
		{BubbleSort, "BubbleSort"},
	}

	for _, sortFunc := range sortFunctions {
		tests := createTestCases()
		for _, test := range tests {
			rawData := make([]int, len(test.data))
			copy(rawData, test.data)

			actualResult := sortFunc.sort(test.data)

			if !equals(actualResult, test.expected) {
				t.Errorf("%s failed on %v: actualResult is %v but expected %v",
					sortFunc.name, rawData, actualResult, test.expected)
			}
		}
	}
}

func createTestCases() []testCase {
	return []testCase{
		{[]int{1, 9, 0, 2, 8, 4, 3, 7, 6, 5}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{[]int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9}},
		{[]int{-1, 10, 345, -99, 69, -1, 25}, []int{-99, -1, -1, 10, 25, 69, 345}},
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
	}
}

type testCase struct {
	data, expected []int
}

func equals(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

package search

import "testing"

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		data          []int
		value         int
		expectedIndex int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 5, 4},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 1, 0},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 9, 8},
		{[]int{9, 9, 9, 9, 9, 9, 9, 9, 9}, 9, 0},
		{[]int{1, 1, 1, 1, 1, 9, 9, 9, 9}, 3, 5},
		{[]int{1, 1, 1, 1, 1}, 100, -1},
		{[]int{1}, 1, 0},
		{[]int{}, 1, -1},
	}

	for _, test := range tests {
		actualIndex := BinarySearch(test.data, test.value)

		if actualIndex != test.expectedIndex {
			t.Errorf("BinarySearch failed on %v: actual found index is %v but expected %v",
				test.data, actualIndex, test.expectedIndex)
		}
	}
}

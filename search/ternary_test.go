package search

import "testing"

func TestTernarySearch(t *testing.T) {
	tests := []struct {
		data          []int
		expectedIndex int
	}{
		{[]int{5, 4, 3, 2, 1, 0, 1, 2, 3, 4}, 5},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, 0},
		{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, 9},
		{[]int{1, 1, 1, 1, 0, 1, 1, 1, 1, 1}, 4},
		{[]int{9}, 0},
		{[]int{}, -1},
	}

	for _, test := range tests {
		actualIndex := TernarySearch(test.data)

		if actualIndex != test.expectedIndex {
			t.Errorf("TernarySearch failed on %v: actual found index is %v but expected %v",
				test.data, actualIndex, test.expectedIndex)
		}
	}
}

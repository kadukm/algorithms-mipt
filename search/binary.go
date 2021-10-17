package search

// BinarySearch searches index of the first element that >= value; data must be sorted ascending.
// It returns -1 if element is not found
func BinarySearch(data []int, value int) int {
	l, r := -1, len(data)-1
	for l+1 < r {
		m := (l + r) / 2
		if data[m] < value {
			l = m
		} else {
			r = m
		}
	}

	if r >= 0 && data[r] >= value {
		return r
	}

	return -1
}

package search

// TernarySearch searches for minimum of data; data should have only one local extrema - its minimum.
// If data is empty then -1 will be returned
func TernarySearch(data []int) int {
	l, r := 0, len(data)-1
	for l+2 < r {
		m1, m2 := (2*l+r)/3, (l+2*r)/3

		if data[m1] > data[m2] {
			l = m1
		} else {
			r = m2
		}
	}

	return l + indexOfMin(data[l:r+1])
}

func indexOfMin(data []int) int {
	if len(data) == 0 {
		return -1
	}

	idx := 0
	min := data[0]
	for i := 1; i < len(data); i++ {
		if data[i] < min {
			idx = i
			min = data[i]
		}
	}

	return idx
}

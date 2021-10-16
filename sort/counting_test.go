package sort

import "testing"

func TestCountingSort(t *testing.T) {
	tests := []struct {
		data, expected []Person
	}{
		{
			[]Person{{1900, "A"}, {2099, "Z"}, {1998, "M"}},
			[]Person{{1900, "A"}, {1998, "M"}, {2099, "Z"}},
		},
		{
			[]Person{{1900, "A"}, {1998, "M"}, {1998, "A"}},
			[]Person{{1900, "A"}, {1998, "M"}, {1998, "A"}},
		},
		{
			[]Person{{1998, "M"}, {1998, "M"}, {1998, "M"}},
			[]Person{{1998, "M"}, {1998, "M"}, {1998, "M"}},
		},
		{
			[]Person{{1998, "M"}},
			[]Person{{1998, "M"}},
		},
		{
			[]Person{},
			[]Person{},
		},
	}

	for _, test := range tests {
		rawData := make([]Person, len(test.data))
		copy(rawData, test.data)

		actualResult := CountingSort(test.data)

		if !personSlicesEqual(actualResult, test.expected) {
			t.Errorf("CountingSort failed on %v: actualResult is %v but expected %v",
				rawData, actualResult, test.expected)
		}
	}
}

func personSlicesEqual(a, b []Person) bool {
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

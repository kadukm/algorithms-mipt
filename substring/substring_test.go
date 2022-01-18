package substring

import "testing"

type SubstringSearch struct {
	apply func(text, substring string) int
	name  string
}

var algorithms = []SubstringSearch{
	{ZFunctionSubstringSearch, "ZFunction"},
}

func TestSubstringSearchAlgorithms(t *testing.T) {
	testCases := []struct {
		text          string
		substring     string
		expectedIndex int
	}{
		{"a", "a", 0},
		{"aa", "a", 0},
		{"ba", "a", 1},
		{"abcdefgh", "def", 3},
		{"abcdefghdefadf", "def", 3},
		{"aaaaaaaaaaaaa", "aaa", 0},
		{"something_abc", "abc", 10},
		{"something_abc", "abc", 10},
		{"", "abc", -1},
		{"asdfas", "", 0},
		{"a", "wrong", -1},
		{"hello my friend", "muy", -1},
		{"abc", "cb", -1},
	}

	for _, algorithm := range algorithms {
		for _, testCase := range testCases {
			result := algorithm.apply(testCase.text, testCase.substring)

			if result != testCase.expectedIndex {
				t.Fatalf(`%s: text="%s", substring="%s", expected %d but was %d`,
					algorithm.name, testCase.text, testCase.substring, testCase.expectedIndex, result)
			}
		}
	}
}

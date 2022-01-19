package trie

import "testing"

func TestTrie(t *testing.T) {
	testCases := []struct {
		strings []string
	}{
		{[]string{"something"}},
		{[]string{"a", "b", "c", "d", "e", "f", "g"}},
		{[]string{"a", "ab", "abc", "abcdefgh"}},
		{[]string{"hello", "world", "and", "my", "dear", "friend"}},
		{[]string{"abcd", "ab", "bcd", "dbcefgh"}},
		{[]string{"a", "a"}},
	}

	for _, testCase := range testCases {
		currTrie := NewTrie(testCase.strings)

		for _, s := range testCase.strings {
			contains := currTrie.Contains(s)
			if !contains {
				t.Fatalf(`expected that the trie contains "%s" but it doesn't'`, s)
			}

			contains = currTrie.Contains("")
			if !contains {
				t.Fatal("expected that the trie contains empty string but it doesn't'")
			}

			contains = currTrie.Contains("wrong string")
			if contains {
				t.Fatalf(`expected that the trie doesn't contain "%s" but it does'`, s)
			}
		}
	}
}

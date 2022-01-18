package substring

// KnuthMorrisPrattAlgorithm searches first index of substring in text or 0 if text doesn't contain substring.
// It cannot work correctly with text containing symbol '#'
func KnuthMorrisPrattAlgorithm(text, substring string) int {
	if len(text) == 0 || len(substring) > len(text) {
		return -1
	}
	if len(substring) == 0 {
		return 0
	}

	p := createPrefixFunction(substring + "#")

	currP := 0
	for i := 0; i < len(text); i++ {
		for currP > 0 && text[i] != substring[currP] {
			currP = p[currP-1]
		}
		if text[i] == substring[currP] {
			currP++
		}

		if currP == len(substring) {
			return i - len(substring) + 1
		}
	}

	return -1
}

func createPrefixFunction(s string) []int {
	p := make([]int, len(s))

	for i := 1; i < len(s); i++ {
		p[i] = p[i-1]
		for p[i] > 0 && s[i] != s[p[i]] {
			p[i] = p[p[i]-1]
		}
		if s[i] == s[p[i]] {
			p[i]++
		}
	}

	return p
}

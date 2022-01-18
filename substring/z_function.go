package substring

// ZFunctionSubstringSearch searches first index of substring in text or 0 if text doesn't contain substring
func ZFunctionSubstringSearch(text, substring string) int {
	if len(text) == 0 || len(substring) > len(text) {
		return -1
	}
	if len(substring) == 0 {
		return 0
	}

	s := substring + text
	z := createZFunction(s)

	for i := len(substring); i < len(s); i++ {
		if z[i] >= len(substring) {
			return i - len(substring)
		}
	}

	return -1
}

func createZFunction(s string) []int {
	z := make([]int, len(s))
	l, r := 0, -1
	for i := 1; i < len(s); i++ {
		if i <= r {
			z[i] = min(r-i+1, z[i-l])
		}
		for i+z[i] < len(s) && s[z[i]] == s[i+z[i]] {
			z[i]++
		}
		if i+z[i]-1 > r {
			l = i
			r = i + z[i] - 1
		}
	}

	return z
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

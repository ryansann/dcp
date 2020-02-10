package test

// match looks for k in n, it returns the index where the first occurrence of k starts in n,
// and the number of iterations it took to find that index, which can be used to validate runtime
// its runtime is O(n + k) where n and k are input sizes
func match(n, k string) (int, int) {
	if len(n) == 0 || len(k) == 0 {
		return -1, 0
	}

	nRunes := []rune(n)
	kRunes := []rune(k)
	st := kRunes[0]

	iterations := 0 // validate runtime
	for i := 0; i < len(nRunes); i++ {
		iterations++
		if nRunes[i] == st { // potential for a match
			found, nPos := true, i
			for j := 1; j < len(kRunes); j++ {
				iterations++
				idx := i + j // index in n
				if idx > len(nRunes)-1 {
					return -1, iterations
				}
				if kRunes[j] != nRunes[idx] {
					found = false
					break
				}
			}
			if found {
				return nPos, iterations
			}
		}
	}

	return -1, iterations
}

package main

// brute force
func strStr(haystack, needle string) int {
	hb, nb := []byte(haystack), []byte(needle)
	if len(hb) < len(nb) {
		return -1
	}

	for i := range hb {
		n1, n2 := i, 0
		for hb[n1] == nb[n2] {
			if n2 == len(nb)-1 {
				return i
			}

			if n1 == len(hb)-1 {
				break
			}
			n1++
			n2++
		}
	}
	return -1
}

// use sliding window
func strStr2(haystack, needle string) int {
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}

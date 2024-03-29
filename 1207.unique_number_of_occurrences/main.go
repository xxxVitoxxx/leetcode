package main

// hash
// time complexity: O(n)
// space complexity: O(n)
func uniqueOccurrences(arr []int) bool {
	occ := make(map[int]int, len(arr))
	for i := 0; i < len(arr); i++ {
		occ[arr[i]]++
	}

	freq := make(map[int]bool, len(occ))
	for _, v := range occ {
		if freq[v] {
			return false
		}
		freq[v] = true
	}

	return true
}

// count occurrences
// time complexity: O(n)
// space complexity: O(n)
func uniqueOccurrences2(arr []int) bool {
	occ := [2001]int{}
	for i := 0; i < len(arr); i++ {
		occ[arr[i]+1000]++
	}

	freq := [2001]int{}
	for i := 0; i < len(occ); i++ {
		if occ[i] > 0 {
			freq[occ[i]]++
		}
	}

	for i := 0; i < len(freq); i++ {
		if freq[i] > 1 {
			return false
		}
	}

	return true
}

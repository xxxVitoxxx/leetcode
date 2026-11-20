package main

// Brute Force
func checkInclusion(s1 string, s2 string) bool {
	var count [26]int
	for i := range s1 {
		count[s1[i]-'a']++
	}

	for i := 0; i < len(s2); i++ {
		remain := count
		j := i
		for j < len(s2) && remain[s2[j]-'a'] > 0 {
			remain[s2[j]-'a']--
			j++
		}

		if j-i == len(s1) {
			return true
		}
	}

	return false
}

// Sliding Window
// N is the length of the s2
// time complexity: O(N)
// time complexity: O(1)
func checkInclusion2(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	var need, window [26]int
	for i := range s1 {
		need[s1[i]-'a']++
		window[s2[i]-'a']++
	}

	if need == window {
		return true
	}

	for i := len(s1); i < len(s2); i++ {
		window[s2[i]-'a']++
		window[s2[i-len(s1)]-'a']--

		if need == window {
			return true
		}
	}

	return false
}

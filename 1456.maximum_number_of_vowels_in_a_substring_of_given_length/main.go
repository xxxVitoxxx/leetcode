package main

// use sliding window
// time complexity: O(n)
// space complexity: O(1)
func maxVowels(s string, k int) int {
	ans, count := 0, 0
	for i := 0; i < len(s); i++ {
		if isVowel(s[i]) {
			count++
		}
		if i >= k && isVowel(s[i-k]) {
			count--
		}
		if count > ans {
			ans = count
		}
	}
	return ans
}

func isVowel(b byte) bool {
	switch b {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	}
	return false
}

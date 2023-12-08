package main

func lengthOfLongestSubstring(s string) int {
	var ans int
	for i := 0; i < len(s); i++ {
		m := make(map[byte]struct{})
		var length int
		for j := i; j < len(s); j++ {
			if _, ok := m[s[j]]; ok {
				break
			}

			m[s[j]] = struct{}{}
			length++
		}

		if length > ans {
			ans = length
		}
	}
	return ans
}

// use sliding window
func lengthOfLongestSubstring2(s string) int {
	m := make(map[byte]int)
	var ans int
	left, right := 0, 0
	for right < len(s) {
		m[s[right]]++
		for m[s[right]] > 1 {
			m[s[left]]--
			left++
		}

		if right-left+1 > ans {
			ans = right - left + 1
		}
		right++
	}
	return ans
}

// sliding window optimized
// time complexity: O(n)
// space complexity: O(min(n, m)) m is the size of the actual character set in the string.
func lengthOfLongestSubstring3(s string) int {
	var ans int
	m := make(map[byte]int)
	for left, right := 0, 0; right < len(s); right++ {
		if idx, ok := m[s[right]]; ok {
			left = max(left, idx)
		}

		ans = max(ans, right-left+1)
		m[s[right]] = right + 1
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// sliding window optimized
// time complexity: O(n)
// space complexity: O(1) the character set is fixed.
func lengthOfLongestSubstring4(s string) int {
	var ans int
	// 7 位版本的 ASCII 有 128 個字符
	// 8 位版本的有 256 個字符
	dict := make([]int, 128)
	for left, right := 0, 0; right < len(s); right++ {
		if left < dict[s[right]] {
			left = dict[s[right]]
		}

		if ans < right-left+1 {
			ans = right - left + 1
		}

		dict[s[right]] = right + 1
	}
	return ans
}

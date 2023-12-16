package main

func lengthOfLongestSubstringTwoDistinct(s string) int {
	var ans int
	for left := 0; left < len(s); left++ {
		m := make(map[byte]int)
		right := left
		for right < len(s) && len(m) < 3 {
			if _, ok := m[s[right]]; !ok && len(m) == 2 {
				break
			}
			m[s[right]]++
			right++
		}

		if right-left > ans {
			ans = right - left
		}
	}
	return ans
}

// use sliding window
func lengthOfLongestSubstringTwoDistinct2(s string) int {
	var ans int
	m := make(map[uint8]int)
	left, right := 0, 0
	for ; right < len(s); right++ {
		m[s[right]] = right
		if len(m) == 3 {
			idx := getMinIndex(m)
			delete(m, s[idx])
			left = idx + 1
		}

		if right-left+1 > ans {
			ans = right - left + 1
		}
	}
	return ans
}

func getMinIndex(m map[uint8]int) int {
	min := 1<<63 - 1
	for _, val := range m {
		if val < min {
			min = val
		}
	}
	return min
}

func lengthOfLongestSubstringTwoDistinct3(s string) int {
	var ans int
	count := make([]int, 128)
	left, right, dist := 0, 0, 0
	for ; right < len(s); right++ {
		count[s[right]]++
		if count[s[right]] == 1 {
			dist++
		}

		for dist > 2 {
			count[s[left]]--
			if count[s[left]] == 0 {
				dist--
			}
			left++
		}

		if right-left+1 > ans {
			ans = right - left + 1
		}
	}
	return ans
}

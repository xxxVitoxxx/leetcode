package main

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	var ans int
	m := make(map[uint8]int)
	var pnt1 int
	for pnt2 := 0; pnt2 < len(s); pnt2++ {
		m[s[pnt2]]++
		for len(m) > k {
			m[s[pnt1]]--
			if m[s[pnt1]] == 0 {
				delete(m, s[pnt1])
			}
			pnt1++
		}

		if pnt2-pnt1+1 > ans {
			ans = pnt2 - pnt1 + 1
		}
	}
	return ans
}

func lengthOfLongestSubstringKDistinct2(s string, k int) int {
	var ans int
	m := make(map[uint8]int)
	var left int
	for right := 0; right < len(s); right++ {
		m[s[right]] = right
		if len(m) > k {
			minIdx := getMinIndex(m)
			delete(m, s[minIdx])
			left = minIdx + 1
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

func lengthOfLongestSubstringKDistinct3(s string, k int) int {
	var ans int
	count := make([]int, 128)
	left, right, dist := 0, 0, 0
	for ; right < len(s); right++ {
		count[s[right]]++
		if count[s[right]] == 1 {
			dist++
		}

		for dist > k {
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

func lengthOfLongestSubstringKDistinct4(s string, k int) int {
	var maxLen int
	m := make(map[uint8]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
		if len(m) <= k {
			maxLen++
		} else {
			m[s[i-maxLen]]--
			if m[s[i-maxLen]] == 0 {
				delete(m, s[i-maxLen])
			}
		}
	}
	return maxLen
}

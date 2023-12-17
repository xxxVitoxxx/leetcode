package main

// time complexity: O(26N)
func longestSubstring(s string, k int) int {
	var ans int
	maxUnique := getMaxUniqueLetters(s)
	for dist := 1; dist <= maxUnique; dist++ {
		ans = max(ans, helper(s, k, dist))
	}
	return ans
}

func getMaxUniqueLetters(s string) int {
	unique, counter := 0, make(map[uint8]int)
	for i := 0; i < len(s); i++ {
		if counter[s[i]] == 0 {
			unique++
		}
		counter[s[i]]++
	}
	return unique
}

func helper(s string, k, dist int) int {
	var res int
	var j int
	var d int // the number of characters where the frequency >= k
	counter := make(map[uint8]int)
	for i := 0; i < len(s); i++ {
		for j < len(s) && len(counter) <= dist {
			counter[s[j]]++
			if counter[s[j]] == k {
				d++
			}

			if len(counter) == dist && d == dist {
				res = max(res, j-i+1)
			}
			j++
		}

		counter[s[i]]--
		if counter[s[i]] == k-1 {
			d--
		}
		if counter[s[i]] == 0 {
			delete(counter, s[i])
		}
	}
	return res
}

func longestSubstring2(s string, k int) int {
	var ans int
	maxUnique := getMaxUniqueLetters(s)
	for currUnique := 1; currUnique <= maxUnique; currUnique++ {
		counter := make([]int, 26)
		start, end := 0, 0
		unique, countAtLeastK := 0, 0
		for end < len(s) {
			if unique <= currUnique {
				counter[s[end]-'a']++
				if counter[s[end]-'a'] == 1 {
					unique++
				}
				if counter[s[end]-'a'] == k {
					countAtLeastK++
				}
				end++
			} else {
				counter[s[start]-'a']--
				if counter[s[start]-'a'] == k-1 {
					countAtLeastK--
				}
				if counter[s[start]-'a'] == 0 {
					unique--
				}
				start++
			}

			if unique == currUnique && unique == countAtLeastK {
				ans = max(ans, end-start)
			}
		}
	}
	return ans
}

// use recursion
// time complexity: O(N)
// [ O {X X X} O {X X}]
func longestSubstring3(s string, k int) int {
	counter := make(map[uint8]int)
	for i := 0; i < len(s); i++ {
		counter[s[i]]++
	}

	var flag bool
	for _, c := range counter {
		if c < k {
			flag = true
			break
		}
	}
	if !flag {
		return len(s)
	}

	var res int
	for i := 0; i < len(s); i++ {
		if counter[s[i]] < k {
			continue
		}

		j := i
		for j < len(s) && counter[s[j]] >= k {
			j++
		}
		res = max(res, longestSubstring3(s[i:j], k))
		i = j
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// use "divide and conquer" and recursion
func longestSubstring4(s string, k int) int {
	return helper4(s, 0, len(s), k)
}

func helper4(s string, start, end, k int) int {
	counter := make([]int, 26)
	for i := start; i < end; i++ {
		counter[s[i]-'a']++
	}

	for mid := start; mid < end; mid++ {
		if counter[s[mid]-'a'] >= k {
			continue
		}

		midNext := mid + 1
		for midNext < end && counter[s[midNext]-'a'] < k {
			midNext++
		}

		return max(helper4(s, start, mid, k), helper4(s, midNext, end, k))
	}
	return end - start
}

package main

func shortestToChar(s string, c byte) []int {
	res := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			var dist int
			res[i] = dist
			left, right := i-1, i+1
			for left >= 0 && s[left] != c || right < len(s) && s[right] != c {
				dist++
				if left >= 0 && s[left] != c {
					if res[left] == 0 {
						res[left] = dist
					} else {
						res[left] = min(res[left], dist)
					}
					left--
				}
				if right < len(s) && s[right] != c {
					if res[right] == 0 {
						res[right] = dist
					} else {
						res[right] = min(res[right], dist)
					}
					right++
				}
			}
		}
	}
	return res
}

func shortestToChar2(s string, c byte) []int {
	res := make([]int, len(s))
	for dist, i := -len(s), 0; i < len(s); i++ {
		if s[i] == c {
			dist = 0
		} else {
			dist++
		}
		res[i] = dist
	}

	for dist, i := 0, len(s)-1; i >= 0; i-- {
		if s[i] == c {
			dist = 1
		} else if dist > 0 {
			if res[i] < 0 || res[i] > dist {
				res[i] = dist
				dist++
			} else {
				dist = 0
			}
		}
	}
	return res
}

func shortestToChar3(s string, c byte) []int {
	res := make([]int, len(s))
	firstOccurrence, lastOccurrence := -1, 0
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			if firstOccurrence == -1 {
				firstOccurrence = i
			}
			lastOccurrence = i
		}
	}

	for i := 0; i < len(s); i++ {
		if s[i] == c {
			res[i] = 0
			firstOccurrence = i
		} else {
			res[i] = abs(i - firstOccurrence)
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == c {
			lastOccurrence = i
		} else {
			res[i] = min(res[i], abs(i-lastOccurrence))
		}
	}
	return res
}

func shortestToChar4(s string, c byte) []int {
	res := make([]int, len(s))
	for i := range res {
		// s.length <= 10^4
		res[i] = 10001
	}

	occurrence := -1
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			res[i] = 0
			occurrence = i
		} else if occurrence != -1 {
			res[i] = min(res[i], abs(i-occurrence))
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == c {
			occurrence = i
		} else if occurrence != -1 {
			res[i] = min(res[i], abs(i-occurrence))
		}
	}
	return res
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func min(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}

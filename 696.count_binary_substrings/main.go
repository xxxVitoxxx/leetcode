package main

// brute force
func countBinarySubstrings(s string) int {
	var res int
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && s[i] == '0' {
			if s[i+1] == '1' {
				res++
			} else {
				idx, l := i, 0
				for ; idx < len(s) && s[idx] == '0'; idx++ {
					l++
				}

				for ; idx < len(s) && s[idx] == '1'; idx++ {
					l--
				}

				if l <= 0 {
					res++
				}
			}
		} else if s[i] == '1' && i+1 < len(s) {
			if s[i+1] == '0' {
				res++
			} else {
				idx, l := i, 0
				for ; idx < len(s) && s[idx] == '1'; idx++ {
					l++
				}

				for ; idx < len(s) && s[idx] == '0'; idx++ {
					l--
				}

				if l <= 0 {
					res++
				}
			}
		}
	}
	return res
}

// use two pointer
func countBinarySubstrings2(s string) int {
	var count int
	check := func(left, right int) {
		currentCount := 1
		for left > 0 && right < len(s)-1 {
			if s[left-1] == s[left] && s[right] == s[right+1] {
				currentCount++
				left, right = left-1, right+1
			} else {
				break
			}
		}
		count += currentCount
	}
	for i := 0; i < len(s)-1; i++ {
		if s[i] != s[i+1] {
			check(i, i+1)
		}
	}
	return count
}

func countBinarySubstrings3(s string) int {
	count, countZero, countOne := 0, 0, 0
	prev := rune(s[0])
	for _, r := range s {
		if r == prev {
			if r == '0' {
				countZero++
			} else {
				countOne++
			}
		} else {
			count += min(countZero, countOne)
			if r == '0' {
				countZero = 1
			} else {
				countOne = 1
			}
		}
		prev = r
	}
	return count + min(countZero, countOne)
}

func countBinarySubstrings4(s string) int {
	count, first, second := 0, 1, 0
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			first++
		} else {
			count += min(first, second)
			second, first = first, 1
		}
	}
	return count + min(first, second)
}

func min(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}

func countBinarySubstrings5(s string) int {
	count, curr, prev := 0, 1, 0
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			curr++
		} else {
			prev, curr = curr, 1
		}
		if prev >= curr {
			count++
		}
	}
	return count
}

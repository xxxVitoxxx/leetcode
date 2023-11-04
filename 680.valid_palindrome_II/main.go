package main

func validPalindrome(s string) bool {
	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
		if s[l] != s[r] {
			return isPalindrome(l+1, r, s) || isPalindrome(l, r-1, s)
		}
	}
	return true
}

func isPalindrome(l, r int, s string) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l, r = l+1, r-1
	}
	return true
}

func validPalindrome2(s string) bool {
	var lSkip bool
	var rSkip bool
	var history [2]int
	for l, r := 0, len(s)-1; l < r; {
		if s[l] == s[r] {
			l, r = l+1, r-1
			continue
		} else if !lSkip {
			lSkip = true
			history[0], history[1] = l, r
			l++
			continue
		} else if !rSkip {
			rSkip = true
			l = history[0]
			r = history[1] - 1
			continue
		}
		return false
	}
	return true
}

func validPalindrome3(s string) bool {
	l, r := 0, len(s)-1
	left, right, valid := isPalindrome3(l, r, s)
	if valid {
		return valid
	}

	if _, _, valid := isPalindrome3(left+1, right, s); valid {
		return valid
	}

	if _, _, valid := isPalindrome3(left, right-1, s); valid {
		return valid
	}
	return false
}

func isPalindrome3(l, r int, s string) (int, int, bool) {
	for l < r {
		if s[l] != s[r] {
			return l, r, false
		}
		l, r = l+1, r-1
	}
	return l, r, true
}

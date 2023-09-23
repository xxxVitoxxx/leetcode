package main

// follow up: can not converting the integer to a string
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	res, y := 0, x
	for ; y > 0; y /= 10 {
		res = res*10 + (y % 10)
	}
	return res == x
}

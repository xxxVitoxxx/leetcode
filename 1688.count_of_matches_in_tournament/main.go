package main

// use recursion
func numberOfMatches(n int) int {
	if n <= 1 {
		return 0
	}

	var match int
	if n%2 == 0 {
		match = n >> 1
		n = n >> 1
	} else {
		match = (n - 1) >> 1
		n = n - (n-1)>>1
	}
	return match + numberOfMatches(n)
}

func numberOfMatches2(n int) int {
	var match int
	for n != 1 {
		if n%2 == 0 {
			match += n / 2
			n /= 2
		} else {
			match += (n - 1) / 2
			n = n - (n-1)/2
		}
	}
	return match
}

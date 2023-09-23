package main

import "math"

// e.g. 451 = 4*10^2 + 5*10^1 + 1*10^0

func titleToNumber(columnTitle string) int {
	var res int
	for _, r := range columnTitle {
		res = res*26 + calculate(r)
	}
	return res
}

func titleToNumber2(columnTitle string) int {
	l := len(columnTitle) - 1
	var res int
	for i, r := range columnTitle {
		c := calculate(r)
		if i == l {
			res += c
		} else {
			cal := c * int(math.Pow(26, float64(l-i)))
			res += cal
		}
	}
	return res
}

func calculate(r rune) int {
	return int(r-'A') + 1
}

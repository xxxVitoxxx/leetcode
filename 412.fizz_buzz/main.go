package main

import "strconv"

func fizzBuzz(n int) []string {
	result := make([]string, n)
	var answer string
	for i := 1; i <= n; i++ {
		switch {
		case i%15 == 0:
			answer = "FizzBuzz"
		case i%3 == 0:
			answer = "Fizz"
		case i%5 == 0:
			answer = "Buzz"
		default:
			answer = strconv.Itoa(i)
		}
		result[i-1] = answer
	}
	return result
}

func fizzBuzz2(n int) []string {
	result := make([]string, n)
	for i := 1; i <= n; i++ {
		switch i % 15 {
		case 0:
			result[i-1] = "FizzBuzz"
		case 3, 6, 9, 12:
			result[i-1] = "Fizz"
		case 5, 10:
			result[i-1] = "Buzz"
		default:
			result[i-1] = strconv.Itoa(i)
		}
	}
	return result
}

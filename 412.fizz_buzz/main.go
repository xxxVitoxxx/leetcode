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

// string concatenation and we can put all mappings in a hash table
func fizzBuzz3(n int) []string {
	result := make([]string, 0, n)

	divisors := []int{3, 5, 7}
	m := map[int]string{
		3: "Fizz",
		5: "Buzz",
		7: "Jazz",
	}

	for i := 1; i <= n; i++ {
		var str string
		for _, divisor := range divisors {
			if i%divisor == 0 {
				str += m[divisor]
			}
		}

		if str == "" {
			str = strconv.Itoa(i)
		}

		result = append(result, str)
	}

	return result
}

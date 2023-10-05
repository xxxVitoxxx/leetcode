package main

func cal(n int) int {
	var sum int
	for ; n > 0; n /= 10 {
		sum = sum + (n%10)*(n%10)
	}
	return sum
}

// use hash table
func isHappy(n int) bool {
	m := make(map[int]struct{})
	for n != 1 {
		if _, ok := m[n]; ok {
			return false
		}

		m[n] = struct{}{}
		n = cal(n)
	}
	return n == 1
}

// use recursive
// when you calculate the sum of the squares of its digits from 0 to 9,
// you will eventually find that only 1 and 7 result in 1.
func isHappy2(n int) bool {
	switch n {
	case 1, 7:
		return true
	case 0, 2, 3, 4, 5, 6, 8, 9:
		return false
	default:
		return isHappy2(cal(n))
	}
}

/*
3
3^2 = 9
9^2 = 81
8^2 + 1^2 = 65
6^2 + 5^2 = 61
6^2 + 1^2 = 37
3^2 + 7^2 = 58
5^2 + 8^2 = 89
8^2 + 9^2 = 145
1^2 + 4^2 + 5^2 = 42
4^2 + 2^2 = 20
2^2 + 0^2 = 4
4^2 = 16
1^2 + 6^2 = 37
3^2 + 7^2 = 58
5^2 + 8^2 = 89
8^2 + 9^2 = 145
*/
// use fast-slow pointers
// whether or not n is a happy number, it will eventually enter a loop.
// if n is a happy number, all the numbers in the loop will be 1.
func isHappy3(n int) bool {
	fast := cal(cal(n))
	slow := cal(n)
	for fast != slow {
		fast = cal(cal(fast))
		slow = cal(slow)
	}
	return slow == 1
}

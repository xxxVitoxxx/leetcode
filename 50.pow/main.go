package main

// binary exponentiation(recursive)
// time complexity: O(log n). at each recursive call we reduce n by half. thus, it will take overall O(log n) time.
// space complexity: O(log n). the recursive stack can use at most O(log n) space at any time.
//
// 2¹⁰⁰ = (2 * 2)⁵⁰
// 4⁵⁰ = (4 * 4)²⁵
// 16²⁵ = 16 * (16 * 16)¹²
// 16 * 256¹² = 16 * (256 * 256)⁶
// 16 * 65536⁶ = 16 * (65536 * 65536)³
// 16 * 4294967296³ = 16 * 4294967296 * (4294967296)²
// 16 * 4294967296³ = 16 * 4294967296 * (4294967296 * 4294967296)¹
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	if n < 0 {
		return 1 / myPow(x, -n)
	}

	if n%2 == 1 {
		return x * myPow(x*x, (n-1)/2)
	}
	return myPow(x*x, n/2)
}

// binary exponentiation(iterative)
// time complexity: O(log n)
// space complexity: O(1). we don't use any additional space.
//
// we can also convert the same recursive approach to an iterative one using some loops.
func myPow2(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}
	if n < 0 {
		n *= -1
		x = 1 / x
	}

	result := 1.0
	for n > 0 {
		if n%2 == 1 {
			result *= x
			n--
		}

		x *= x
		n /= 2
	}
	return result
}

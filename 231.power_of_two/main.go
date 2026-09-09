package main

import "math"

// time complexity: O(logN)
// space complexity: O(1)
func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}

	for n%2 == 0 {
		n /= 2
	}

	return n == 1
}

func isPowerOfTwo2(n int) bool {
	powerOfTwo := math.Log2(float64(n))
	return n > 0 && powerOfTwo == math.Trunc(powerOfTwo)
}

/*
in bit world, to compute `-n` one has to revert all bits in `n` and then add `1` to result.
n := 4
n  -> 00000100
^n -> 11111011
-n -> 11111100 (^n+1)

n := 8
n  -> 00001000
^n -> 11110111
-n -> 11111000 (^n+1)

n := 9
n  -> 00001001
^n -> 11110110
-n -> 11110111 (^n+1)

let's do `n & -n` to keep the `1-bit` and set all the other bits to `0`.
as discussed above, for the power of two, it would result in `n` itself,
since a power of two contains just one `1-bit`.

other numbers have more than `1-bit` in their binary representation and
hence for then `n & -n` would not bet equal `n` itself.

time complexity: O(1)
space complexity: O(1)
*/
func isPowerOfTwo3(n int) bool {
	return n > 0 && n&(-n) == n
}

/*
a number is a power of two if it can be written as: 1, 2, 4, 8, 16, 32
in binary, these number look like:
1  -> 00000001
2  -> 00000010
4  -> 00000100
8  -> 00001000
16 -> 00010000
32 -> 00100000

0  -> 00000000
1  -> 00000001
3  -> 00000011
7  -> 00000111
15 -> 00001111
31 -> 00011111

- there is exactly one `1` bit in the binary form.
- subtracting `1` flips that `1` to `0` and all lower bits to `1`.
- performing n & (n-1) will result in `0` only for powers of two.

time complexity: O(1)
space complexity: O(1)
*/
func isPowerOfTwo4(n int) bool {
	return n > 0 && n&(n-1) == 0
}

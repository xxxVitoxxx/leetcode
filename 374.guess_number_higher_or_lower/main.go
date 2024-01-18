package main

var pick int

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */
func guess(num int) int {
	if num < pick {
		return 1
	} else if num > pick {
		return -1
	}
	return 0
}

// binary search
// time complexity: O(log(n))
// space complexity: O(1)
func guessNumber(n int) int {
	pick, low, high := 0, 1, n
	for low <= high {
		pick = low + (high-low)/2
		res := guess(pick)
		if res == 0 {
			break
		} else if res == 1 {
			low = pick + 1
		} else {
			high = pick - 1
		}
	}
	return pick
}

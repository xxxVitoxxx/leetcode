package main

// time complexity: O(n)
// space complexity: O(1)
func increasingTriplet(nums []int) bool {
	i, j := 1<<32-1, 1<<32-1
	for _, num := range nums {
		if num <= i {
			i = num
		} else if num <= j {
			j = num
		} else {
			return true
		}
	}
	return false
}

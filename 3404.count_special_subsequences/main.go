package main

// time complexity: O(n^2)
// space complexity: O(n)
//
// from nums[p] * nums[r] = nums[q] * nums[s]
// we can get nums[p] / nums[q] = nums[s] / nums[r]
//
// because p < q < r < s, we can split the subsequence into two parts: (p, q) and (r, s)
// it can reduce the time complexity from O(n^4) to O(n^2)
//
// use map to keep track of the ratios of elements in the first part of
// the subsequence (p, q), allowing for quick lookups when processing the
// latter part (r, s).
//
// outer loop (index q): iterate through the array using an index q, starting
// from the third element and stopping before the last four elements. this index
// represents the current ending element of the first part of the subsequence(p, q)
//
// inner loop (index p): for each q, loop through all elements before it using
// index p, calculate the ratio of nums[p] / nums[q] and store this in the map,
// incrementing the count for each ratio found.
//
// second inner loop (index r and s): set r to q + 2, then iterate with s starting
// from r + 2 to the end of the array. for each s, calculate the ratio of nums[s] to
// nums[r] and check if this ratio exists in the map. if it does, add the count of
// that ratio to the result.
func numberOfSubsequences(nums []int) int64 {
	n := len(nums)
	matches := make(map[float64]int64)
	result := int64(0)
	for q := 2; q < n-4; q++ {
		for p := 0; p < q-1; p++ {
			matches[float64(nums[p])/float64(nums[q])]++
		}

		r := q + 2
		for s := r + 2; s < n; s++ {
			result += matches[float64(nums[s])/float64(nums[r])]
		}
	}
	return result
}

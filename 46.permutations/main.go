package main

// time complexity: O(n * n!)
// finding permutation is a well-studied problem in combinatorics.
// given a set of length n, the number of permutation is n!(n factorial).
// there are n options for the first number, n-1 for the second, ans so on.
// for each of the n! permutations, we need O(n) work to copy curr into the answer.
// this gives us O(n * n!) work.

// use hash to record num that is visited.
func permute(nums []int) [][]int {
	ans := make([][]int, 0)
	backtrack(nums, []int{}, &ans, make(map[int]struct{}))
	return ans
}

func backtrack(nums, curr []int, ans *[][]int, used map[int]struct{}) {
	if len(curr) == len(nums) {
		// make a new slice to avoid change the combination in the answer during back track
		copied := make([]int, len(curr))
		copy(copied, curr)
		*ans = append(*ans, copied)
		return
	}

	for _, num := range nums {
		if _, ok := used[num]; !ok {
			used[num] = struct{}{}
			curr = append(curr, num)
			backtrack(nums, curr, ans, used)
			delete(used, num)
			curr = curr[:len(curr)-1]
		}
	}
}

func permute2(nums []int) [][]int {
	ans := make([][]int, 0)
	backtrack2(nums, []int{}, &ans)
	return ans
}

func backtrack2(nums, curr []int, ans *[][]int) {
	// we known that we found a new combination when we have no elements
	if len(nums) == 0 {
		*ans = append(*ans, curr)
		return
	}

	// for the next iteration we consider all the elements but the current one
	for i, num := range nums {
		// make sure to allocate to a new slice
		numsCopy := make([]int, len(nums))
		copy(numsCopy, nums)
		backtrack2(append(numsCopy[:i], numsCopy[i+1:]...), append(curr, num), ans)
	}
}

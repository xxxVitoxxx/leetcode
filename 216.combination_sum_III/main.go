package main

// backtracking
// time complexity: O(9!*k/(9-k)!)
//   - in a worst case, we have to explore all potential combinations to the very end,
//     i.e. the sum n is a large number (n > 9*9). at the first step, we have 9 choices,
//     while at the second step, we have 8 choices, so on and so forth.
//   - the number of exploration we need to make in the worst case would be P(9, k) = 9!/(9-k)!,
//     assuming that k <= 9. by the way, k cannot be greater than 9, otherwise we cannot have
//     a combination whose digits are all unique.
//   - each exploration takes a constant time to process, except the last step where it takes
//     O(k) time to make a copy of combination.
//   - to sum up, the overall time complexity of the algorithm would be 9!/(9-k)! * O(k) = O(9!*k/(9-k)!).
//
// space complexity: O(k)
//   - during the backtracking, we used a list to keep the current combination,
//     which holds up to k elements, i.e. O(k).
//   - since we employed recursion in the backtracking, we would need some additional space for
//     the function call stack, which could pile up to k consecutive invocations, i.e. O(k).
//   - hence, to sum up, the overall space complexity would be O(k).
//   - note that, we did not take into account the space for the final results in the space complexity.
func combinationSum3(k int, n int) [][]int {
	var combinations [][]int
	var backtrack func(sum, nextNum int, combination []int)
	backtrack = func(sum, nextNum int, combination []int) {
		if sum == n && len(combination) == k {
			tmp := make([]int, k)
			copy(tmp, combination)
			combinations = append(combinations, tmp)
			return
		}
		if sum != n && len(combination) == k {
			return
		}

		for i := nextNum; i <= 9; i++ {
			backtrack(sum+i, i+1, append(combination, i))
		}
	}
	backtrack(0, 1, []int{})

	return combinations
}

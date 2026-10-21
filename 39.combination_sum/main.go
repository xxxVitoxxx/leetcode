package main

func combinationSum(candidates []int, target int) [][]int {
	var combination []int
	var result [][]int
	backtracking(combination, candidates, 0, target, &result)
	return result
}

func backtracking(combination, candidates []int, start, remaining int, result *[][]int) {
	if remaining == 0 {
		newCombination := make([]int, len(combination))
		copy(newCombination, combination)
		*result = append(*result, newCombination)
	} else if remaining < 0 {
		return
	}

	for i := start; i < len(candidates); i++ {
		combination = append(combination, candidates[i])
		backtracking(combination, candidates, i, remaining-candidates[i], result)
		combination = combination[:len(combination)-1]
	}
}

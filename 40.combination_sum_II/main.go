package main

import "slices"

func combinationSum2(candidates []int, target int) [][]int {
	slices.Sort(candidates)

	var result [][]int
	var dfs func([]int, int, int)
	dfs = func(comb []int, start, remaining int) {
		if remaining == 0 {
			result = append(result, slices.Clone(comb))
			return
		}

		for i := start; i < len(candidates); i++ {
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}

			if remaining < candidates[i] {
				break
			}

			comb = append(comb, candidates[i])
			dfs(comb, i+1, remaining-candidates[i])
			comb = comb[:len(comb)-1]
		}
	}

	dfs([]int{}, 0, target)

	return result
}

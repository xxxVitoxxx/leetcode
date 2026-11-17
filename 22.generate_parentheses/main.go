package main

func generateParenthesis(n int) []string {
	var result []string
	var dfs func(string, int, int)
	dfs = func(comb string, left, right int) {
		if len(comb) == 2*n {
			result = append(result, comb)
			return
		}

		if left < n {
			dfs(comb+"(", left+1, right)
		}

		if right < left {
			dfs(comb+")", left, right+1)
		}
	}

	dfs("", 0, 0)

	return result
}

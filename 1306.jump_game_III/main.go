package main

// BFS
// time complexity: O(N)
// space complexity: O(N)
func canReach(arr []int, start int) bool {
	visited := make([]bool, len(arr))
	visited[start] = true
	queue := []int{start}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if arr[curr] == 0 {
			return true
		}

		if curr+arr[curr] < len(arr) && !visited[curr+arr[curr]] {
			queue = append(queue, curr+arr[curr])
			visited[curr+arr[curr]] = true
		}

		if curr-arr[curr] >= 0 && !visited[curr-arr[curr]] {
			queue = append(queue, curr-arr[curr])
			visited[curr-arr[curr]] = true
		}
	}

	return false
}

// DFS
// time complexity: O(N)
// space complexity: O(N)
func canReach2(arr []int, start int) bool {
	visited := make([]bool, len(arr))

	var dfs func(int) bool
	dfs = func(index int) bool {
		if index < 0 || index >= len(arr) || visited[index] {
			return false
		}

		if arr[index] == 0 {
			return true
		}

		visited[index] = true

		return dfs(index+arr[index]) || dfs(index-arr[index])
	}

	return dfs(start)
}

// DFS
// time complexity: O(N)
// space complexity: O(N)
func canReach3(arr []int, start int) bool {
	if start < 0 || start >= len(arr) || arr[start] < 0 {
		return false
	}

	if arr[start] == 0 {
		return true
	}

	arr[start] *= -1

	return canReach3(arr, start+arr[start]) || canReach3(arr, start-arr[start])
}

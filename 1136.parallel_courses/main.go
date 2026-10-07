package main

// BFS
func minimumSemesters(n int, relations [][]int) int {
	adjList := make(map[int][]int)
	inDegree := make([]int, n+1)
	for _, r := range relations {
		adjList[r[0]] = append(adjList[r[0]], r[1])
		inDegree[r[1]]++
	}

	var queue []int
	for i := 1; i < len(inDegree); i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	var semesters, studyCount int
	for len(queue) > 0 {
		size := len(queue)
		semesters++

		for i := range size {
			studyCount++

			for _, next := range adjList[queue[i]] {
				inDegree[next]--
				if inDegree[next] == 0 {
					queue = append(queue, next)
				}
			}
		}

		queue = queue[size:]
	}

	if studyCount != n {
		return -1
	}

	return semesters
}

func minimumSemesters2(n int, relations [][]int) int {
	adjList := make(map[int][]int)
	for _, r := range relations {
		adjList[r[0]] = append(adjList[r[0]], r[1])
	}

	// 0 -> not visited yet
	// 1 -> recursive call in progress
	// 2 -> recursive call finished and no cycle found
	states := make([]int, n+1)
	depth := make([]int, n+1)
	var dfs func(int) bool
	dfs = func(course int) bool {
		if states[course] == 1 {
			return true
		}

		if states[course] == 2 {
			return false
		}

		var longest int
		states[course] = 1
		for _, next := range adjList[course] {
			if dfs(next) {
				return true
			}

			longest = max(longest, depth[next])
		}

		states[course] = 2
		depth[course] = longest + 1

		return false
	}

	var result int
	for course := range adjList {
		if dfs(course) {
			return -1
		}

		result = max(result, depth[course])
	}

	return result
}
